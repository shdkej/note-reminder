FROM python:3.8 as builder
WORKDIR /
COPY ./requirments.txt /
RUN pip3 install -r requirments.txt

FROM python:3.8-slim-buster as recommender
WORKDIR /
COPY --from=builder /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages
COPY ./recommender_system.py /
COPY ./telegram.py /
