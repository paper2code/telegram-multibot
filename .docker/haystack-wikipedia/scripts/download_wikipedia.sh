#!/bin/sh

set -x
set -e

wget -nc -P /opt/data "http://download.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2"
