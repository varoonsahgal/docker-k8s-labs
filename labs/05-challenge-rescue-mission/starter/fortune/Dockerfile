# The "it works, ship it" Dockerfile.
# Your mission (Boss Level): put this image on a serious diet.

FROM python:3.12

COPY . /app
WORKDIR /app

RUN apt-get update
RUN apt-get install -y vim curl build-essential   # "just in case"
RUN pip install -r requirements.txt

EXPOSE 5000

CMD python app.py
