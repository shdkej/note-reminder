FROM python:3 as builder
WORKDIR / 
COPY ./requirments.txt /
RUN pip3 install -r requirments.txt
# if need initial python dependency using this
#RUN pip3 install --no-cache-dir pandas sklearn requests
#RUN pip3 freeze > requirments.txt

FROM python:3.8-slim-buster as recommender
WORKDIR /
COPY --from=builder /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages
#COPY --from=builder /requirments.txt /
COPY ./recommender_system.py /
COPY ./telegram.py /
