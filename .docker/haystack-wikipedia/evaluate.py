#!/usr/bin/env python3

HOST = 'localhost'
PORT = 9200
INDEX_NAME = 'wikipedia_100_stride_50'

from haystack import Finder
from haystack.reader.transformers import TransformersReader
from haystack.utils import print_answers
from haystack.database.elasticsearch import ElasticsearchDocumentStore

document_store = ElasticsearchDocumentStore(host=HOST, port=PORT, username="", password="", index=INDEX_NAME)

from haystack.retriever.sparse import ElasticsearchRetriever
retriever = ElasticsearchRetriever(document_store=document_store)
READER_DiR = "../models/roberta-base-squad2"
READER_DiR = "../models/electra-base-squad2"
reader = TransformersReader(model=READER_DiR, tokenizer=READER_DiR,  use_gpu=0)

import logging
logging.disable(logging.WARNING)
finder = Finder(reader, retriever)
prediction = finder.get_answers(question="who is the father of Arya Stark",
                                top_k_retriever=10,
                                top_k_reader=3)

print_answers(prediction)
