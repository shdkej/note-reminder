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
    result = []
    result.append("Recommending about -- " + item(item_id))
    result.append("-------------------\n")
    for rec in recs:
        result.append(item(rec[1]))
    return result

def getRecommend():
    contentBasedRecommend()
    num = random.randrange(len(ds))
    result_array = recommend(item_id=num, num=5)
    message = setOutput(result_array)
    send_message(message)

def setOutput(content):
    body = list(map(lambda b: b.replace("==", "\n"), content))
    result = list(map(lambda a: "**" + a.split("\n", 1)[0] + "**" + a.split("\n", 1)[-1] + "\n", body))
    result.append("shdkej.com")
    return result


getRecommend()
