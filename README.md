# mockDSPs

これは実はDSPは関係なく、主要な機能としてはJSON Requestに対してtemplateを書くことによって動的にレスポンスを返すことができるだけの鯖である

ただなんとなく一応Nativeのadmも頑張れば返せるようにはしてある

`{{ .payload $imp[0].native.request }}` みたいな感じでtemplateかくとnative objectのrequestとNativeAdm.tmplにそって返す気がする

## for example

hoge.tmpl
``` json
{"reqid": "{{ .id }}" }
```

``` shell
curl -X POST localhost:8080/mockDSPs/hoge -d '{"id":"piyo"}'
```
