#!/bin/sh

set -x
set -e

python3 -m wikiextractor.WikiExtractor -o "/opt/data/wikipedia/" --json \
--filter_disambig_page \
--processes 8 \
"/opt/data/enwiki-latest-pages-articles.xml.bz2"
