import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import linear_kernel
import boto3
import logging
import random
import json

logger = logging.getLogger()
if logger.handlers:
    for handler in logger.handlers:
        logger.removeHandler(handler)
logging.basicConfig(level=logging.INFO)

bucket = 'my-note-0.0.1'
key = 'tags.csv'
obj = '/tmp/tags.csv'
results = {}
ds = None
sqs_client = None

def init_data():
    """Lambda cold start 시 S3에서 데이터 로드"""
    global ds, sqs_client
    if ds is None:
        logger.info(f"Downloading {key} from S3 bucket {bucket}")
        s3 = boto3.resource('s3')
        s3.meta.client.download_file(bucket, key, obj)
        ds = pd.read_csv(obj)
        logger.info(f"Loaded {len(ds)} records from CSV")
    if sqs_client is None:
        sqs_client = boto3.client('sqs', region_name='eu-central-1')

def send_to_sqs(message):
    """SQS로 메시지 발송 (Telegram 발송 Lambda가 처리)"""
    queue_url = 'https://sqs.eu-central-1.amazonaws.com/917213086376/sns-sqs-upload'

    try:
        response = sqs_client.send_message(
            QueueUrl=queue_url,
            MessageBody=json.dumps({'message': message})
        )
        logger.info(f"Message sent to SQS: {response['MessageId']}")
        return True
    except Exception as e:
        logger.error(f"Failed to send message to SQS: {e}")
        return False
def contentBasedRecommend():
    tf = TfidfVectorizer(analyzer='word', ngram_range=(1, 3), min_df=0, stop_words='english')
    tfidf_matrix = tf.fit_transform(ds['description'].values.astype('U'))

    cosine_similarities = linear_kernel(tfidf_matrix, tfidf_matrix)

    for idx, row in ds.iterrows():
        similar_indices = cosine_similarities[idx].argsort()[:-100:-1]
        similar_items = [(cosine_similarities[idx][i], ds['id'][i]) for i in similar_indices]

        results[row['id']] = similar_items[1:]
        
    print('done!')

def item(id):
    return ds.loc[ds['id'] == id]['description'].tolist()[0].split(' - ')[0]

# Just reads the results out of the dictionary.
def recommend(item_id, num):
    recs = results[item_id][:num]
    return {
        'source': item(item_id),
        'recommendations': [(item(rec[1]), rec[0]) for rec in recs]
    }

def getRecommend(event, context):
    # S3에서 데이터 로드 (cold start 시에만)
    init_data()

    contentBasedRecommend()
    number = 1
    if event.get('num'):
        number = event.get('num')
    elif event.get('Message'):
        number = event.get('Message')
    else:
        number = random.randrange(len(ds))

    logger.info(f'Picked Integer: {number}')
    result_data = recommend(item_id=number, num=5)
    message = setOutput(result_data)

    # SQS로 메시지 발송 (Telegram Lambda가 처리)
    send_to_sqs(message)

    return {'message': message}

NUM_EMOJIS = ['1️⃣', '2️⃣', '3️⃣', '4️⃣', '5️⃣']

def setOutput(data):
    source_text = data['source'].replace("==", "\n")
    source_title = source_text.split("\n", 1)[0]
    source_body = source_text.split("\n", 1)[-1] if "\n" in source_text else ""

    lines = []
    lines.append("📚 *추천 노트*\n")
    lines.append(f"📌 *{source_title}*")
    if source_body:
        lines.append(source_body)
    lines.append("\n━━━━━━━━━━\n")

    for i, (text, score) in enumerate(data['recommendations']):
        item_text = text.replace("==", "\n")
        title = item_text.split("\n", 1)[0]
        body = item_text.split("\n", 1)[-1] if "\n" in item_text else ""
        score_pct = int(score * 100)
        emoji = NUM_EMOJIS[i] if i < len(NUM_EMOJIS) else f"{i+1}."

        lines.append(f"{emoji} *{title}* ({score_pct}%)")
        if body:
            lines.append(body)
        lines.append("")

    lines.append("🔗 shdkej.com")
    return "\n".join(lines)

