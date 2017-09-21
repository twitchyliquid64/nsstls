#!/bin/bash

function build
{
  go build -o libnss_tls.so -buildmode c-shared
}

function clean
{
  go clean
  rm -f libnss_tls.h libnss_tls.so
}

function install
{
  export wd=`pwd`
  sudo rm /lib/libnss_tls.so.2
  sudo cp libnss_tls.so /lib/libnss_tls.so.2
  cd /lib
  sudo install -m 0644 libnss_tls.so.2 /lib
  sudo /sbin/ldconfig -n /lib /usr/lib
  cd $wd
}

function uninstall
{
  sudo rm /lib/libnss_tls.so.2
}

function test
{
  pamtester test_pamtls test authenticate
}

case $1 in
  "clean")
    clean
    ;;
  "clean_build")
    clean
    build
    ;;
  "build")
    build
    ;;

  "install")
    install
    ;;
  "uninstall")
    uninstall
    ;;
  "test")
    test
    ;;

  *)
    echo "No argument specified. Defaulting to 'clean_build'."
    echo "In future, invoke like: \n  ${0} clean/clean_build/build/install_test/uninstall_test"
    build
    ;;
esac
