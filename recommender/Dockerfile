FROM python:3 as builder
WORKDIR / 
RUN pip3 install --no-cache-dir pandas sklearn
RUN pip3 freeze > requirments.txt
#COPY ./requirments.txt /
#RUN pip3 install -r requirments.txt

FROM python:3.8-slim-buster
WORKDIR /
COPY --from=builder /usr/local/lib/python3.8/site-packages /usr/local/lib/python3.8/site-packages
COPY --from=builder /requirments.txt /
COPY ./ /
CMD python3 recommender_system.py
