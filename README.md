# IoT API Sample

## 概要

![Outline of this system](./images/outline_of_the_system.png)

これは、以下の特徴をもったシステムのモック実装です。

* Arduinoに接続したセンサーを[Gobotライブラリ](https://gobot.io/documentation/platforms/arduino/)を用いて読み取る
* [gin-gonicライブラリ](https://github.com/gin-gonic/gin)を用いたREST APIサービスを作る
* センサー読み取り値をREST APIで公開する（JSON over HTTP）
