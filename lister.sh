#!/bin/bash

rm all.stat.csv

for folder in internal/*; do
  echo "------ ${folder} ------"
  name=$(echo ${folder} | sed 's@/@@') # remove slash internal/some -> internalsome
  inarray=$(ls | grep -o "${name}Go" | wc -w)
  if test $inarray -eq 0
  then
    echo "run benchmarks for ${folder}"
    make jbench path=./${folder} | tee ${name}Go1.17
    GO_BIN_PATH=~/go/go1.18/bin/go make jbench path=./${folder} | tee ${name}Go1.18
    GO_BIN_PATH=~/go/go1.18/bin/go make gbench path=./${folder} | tee ${name}Go1.18Generics
  else
    echo "${folder} benchmarks exist"
  fi

  sed 's/G\//J\//g' ${name}Go1.18Generics > ${name}Go1.18GenericsFix
  benchstat ${name}Go1.17 ${name}Go1.18 ${name}Go1.18GenericsFix > ${name}.stat.txt
  benchstat -csv ${name}Go1.17 ${name}Go1.18 ${name}Go1.18GenericsFix > ${name}.stat.csv
  benchstat -csv ${name}Go1.17 ${name}Go1.18 ${name}Go1.18GenericsFix >> all.stat.csv
done
