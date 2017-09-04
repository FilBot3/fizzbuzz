#!/usr/bin/env python
#
#

from setuptools import setup

setup(
  name='fizzbuzz',
  version='0.1.0',
  description='Performs the fizzbuzz challenge',
  author='Phillip Dudley',
  author_email='Predatorian3@gmail.com',
  license='MIT',
  packages=['fizzbuzz'],
  install_requires=[],
  tests_require=[
    'nose'
  ],
  test_suite='nose.collector'
)