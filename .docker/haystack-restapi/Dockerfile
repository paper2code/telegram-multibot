FROM python:3.7.4-stretch

WORKDIR /home/user

RUN git clone --depth=1 https://github.com/deepset-ai/haystack

# install as a package
WORKDIR /home/user/haystack

RUN pip3 install -r requirements.txt
RUN pip3 install -e .

EXPOSE 8000

# cmd for running the API
CMD ["gunicorn", "rest_api.application:app",  "-b", "0.0.0.0", "-k", "uvicorn.workers.UvicornWorker", "--workers", "2", "--timeout", "180"]
