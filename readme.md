Introduction
============
Golang sdk integrates cashier, transfer, transaction, betting, airtime module business, users can access opay more quickly

Official document: https://documentation.opayweb.com

cashier
=======
>**H5 Checkout (For Web based payments)**  


+ Use the OPay Checkout to quickly collect payments for purchases on your web platform without having to build a checkout page. It provides a pre-built UI solution that helps you process payments from the available methods.

+ The OPay Checkout API is the most popular way to integrate pay with OPay. It provides a pre-built UI solution for processing payments that handles checkout for you, if you don’t want to build a checkout page.

>**sdk**

+ examples/opaycashier
+ + checkout
  
    Creating a cashier order will return to the h5 payment page. After the user pays the order, it will call back the specified callback and the web() function will process the verification signature  
    
    https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/checkout.go
  + close
  
    only init status order can close  

    https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/close.go

  + refund
    
    only succeed order can refund。 The case contains the refund to the wallet, bank card, bank account  
    
    https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/refund.go

transfer
========

transaction
===========

betting
=======

airtime
=======


