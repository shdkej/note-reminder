import pandas as pd
import random
from sklearn.feature_extraction.text import TfidfVectorizer
from sklearn.metrics.pairwise import linear_kernel
import os
from telegram import send_message

CSV_PATH = os.environ['CSV_PATH']
results = {}
ds = pd.read_csv(CSV_PATH)

def contentBasedRecommend():
    tf = TfidfVectorizer(analyzer='word', ngram_range=(1, 3), min_df=0, stop_words='english')
    tfidf_matrix = tf.fit_transform(ds['description'])

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

def getRecommend():
    contentBasedRecommend()
    num = random.randrange(len(ds))
    result_data = recommend(item_id=num, num=5)
    message = setOutput(result_data)
    send_message(message)

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


getRecommend()
