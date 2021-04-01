Introduction
============
Golang sdk integrates cashier, transfer, transaction, betting, airtime module business, users can access opay more quickly

Official document: https://documentation.opayweb.com  

sdk merchant config [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/multimerchant/cashier.go)
===================
>**support multi merchant config, conf.InitEnv can call multitimes, but one merchant id only init once. **. 


cashier
=======
>**H5 Checkout (For Web based payments)**  


+ Use the OPay Checkout to quickly collect payments for purchases on your web platform without having to build a checkout page. It provides a pre-built UI solution that helps you process payments from the available methods.

+ The OPay Checkout API is the most popular way to integrate pay with OPay. It provides a pre-built UI solution for processing payments that handles checkout for you, if you don’t want to build a checkout page.

>**sdk** [here](https://github.com/opay-services/opay-sdk-golang/blob/master/sdk/cashier)

+ examples/opaycashier
+ + checkout [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/checkout.go)
  
    Creating a cashier order will return to the h5 payment page. After the user pays the order, it will call back the specified callback and the web() function will process the verification signature  
    
  + close [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/close.go)
  
    only init status order can close  

  + refund [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaycashier/refund.go)  
  
    only succeed order can refund。 The case contains the refund to the wallet, bank card, bank account  
    

transfer
========
**transfer is merchant end to transfer money, to bank account, opay owallet。Currently, there is no callback for transfers, and developers can only keep track of the transfer progress through status inquiries**  
-----------
-----------  
**It is recommended to use the wallet to transfer money, with fewer call chains and fast arrival speed, if you select to bank account , the transfer will be successful after a few minutes, depending on the processing time of the bank**
----------
----------

>**sdk**  [here](https://github.com/opay-services/opay-sdk-golang/blob/master/sdk/transfer)

+ examples/opaytransfer  

+ + towallet (tow kinds of merchant and user)  
  
  + + to merchant [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransfer/toWalletMerchant.go)  
    + to user [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransfer/toWalletUser.go)


+ + to bank only support bank account   [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransfer/toWalletUser.go)



transaction
===========
**transaction user-to-merchant transfer, user can user bank card, bank account to pay。we support ussd dail mode。**  
----------
---------
**also support merchant support a bank account, the user recharges to a given account through other channels  **  
----------  
----------
>**sdk** [here](https://github.com/opay-services/opay-sdk-golang/blob/master/sdk/transaction)

**any mode,  after success, you will receive callback by specify url ** 

+ examples/opaytransaction  
* * user by bank account  [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransaction/bybankaccount.go) 

* * user by bank card [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransaction/bybankcard.go)  

* * ussd dail [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransaction/byussd.go)  

* * merchant support a bank account, the user recharges to a given account through other channels [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/opaytransaction/banktransfer.go)



betting
=======
>**sdk** [here](https://github.com/opay-services/opay-sdk-golang/blob/master/sdk/betting)

+ examples/betting [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/betting/bet.go)

airtime
=======
>**sdk** [here](https://github.com/opay-services/opay-sdk-golang/blob/master/sdk/airtime)  

+ examples/betting [here](https://github.com/opay-services/opay-sdk-golang/blob/master/examples/airtime/topup.go)



