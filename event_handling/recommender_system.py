import pandas as pd
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import linear_kernel
import boto3
import logging
import random

logger = logging.getLogger()
if logger.handlers:
    for handler in logger.handlers:
        logger.removeHandler(handler)
logging.basicConfig(level=logging.INFO)

bucket = 'my-note-0.0.1'
key = 'tags.csv'
s3 = boto3.resource('s3')
s3.meta.client.download_file(bucket, key, '/tmp/tags.csv')
obj = '/tmp/tags.csv'
results = {}
ds = pd.read_csv(obj)
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
    print("Recommending " + str(num) + " products similar to " + item(item_id) + "...")
    print("-------")
    recs = results[item_id][:num]
    result = []
    result.append("Recommending about -- " + item(item_id))
    result.append("-------------------\n")
    for rec in recs:
        print("Recommended: " + item(rec[1]) + " (score:" + str(rec[0]) + ")")
        result.append(item(rec[1]))
    return result

def getRecommend(event, context):
    contentBasedRecommend()
    number = 1
    if event.get('num'):
        number = event.get('num')
    elif event.get('Message'):
        number = event.get('Message')
    else:
        number = random.randrange(len(ds))

    # logger.info('Picked Integer: {}'.format(num))
    result_array = recommend(item_id=number, num=5)
    message = setOutput(result_array)
    return {'message': message}

def setOutput(content):
    body = list(map(lambda b: b.replace("==", "\n"), content))
    result = list(map(lambda a: "*" + a.split("\n", 1)[0] + "*\n" + a.split("\n", 1)[-1], body))
    result.append("wasm.shdkej.com")
    return result

