<template>
	<h2>微信扫码支付</h2>
	<img src="/images/market/pay.jpg" style="margin-left: 300px; height: 400px; width: 300px;">
</template>




<script>
import QRCode from 'qrcodejs2'

export default {
	 data() {
	   		return { }
	 },
	  methods:{
			getCoed() {
		        let qrcode = new QRCode('qrcode', {
		        width: 200,  // 宽
		        height: 200, // 高
		        text: 'weixin://wxpay/bizpayurl?pr=X8D912ad', // 二维码内容,也可以直接放接口获取的数据
		        // render: 'canvas'     // 设置渲染方式（有两种方式 table和canvas，默认是canvas）
		        // background: '#f0f'   // 自定义颜色默认黑白
		        // foreground: '#ff0'
		        })
			},
	  }
}
</script>
