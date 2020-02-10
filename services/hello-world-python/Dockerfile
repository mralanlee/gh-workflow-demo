FROM python:3

WORKDIR /app

ADD main.py requirements.txt /app/

RUN pip install -r requirements.txt

ENTRYPOINT ["python", "main.py"]
