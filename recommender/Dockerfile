FROM python:3 as builder
WORKDIR / 
# if need initial python dependency using this
# RUN pip3 install --no-cache-dir pandas sklearn
# RUN pip3 freeze > requirments.txt
COPY ./requirments.txt /
RUN pip3 install -r requirments.txt

FROM python:3.8-slim-buster as recommender
WORKDIR /
COPY --from=builder /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages
#COPY --from=builder /requirments.txt /
COPY ./recommender_system.py /
COPY ./tags.csv /
RUN python3 recommender_system.py > recommend.txt

FROM python:3.8-slim-buster
WORKDIR /
RUN pip3 install requests \
        python-dotenv
COPY --from=recommender /recommend.txt /
COPY ./ /
RUN python3 telegram.py
