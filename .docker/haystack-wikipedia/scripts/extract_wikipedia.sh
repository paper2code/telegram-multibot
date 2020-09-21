#!/bin/sh

set -x
set -e

python3 -m wikiextractor.WikiExtractor -o "../../shared/datasets/wikipedia/" --json \
--filter_disambig_page \
--processes 8 \
"../../shared/datasets/enwiki-latest-pages-articles.xml.bz2"
