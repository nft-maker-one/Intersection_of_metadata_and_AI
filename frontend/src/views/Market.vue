<template>
    <div class="marketList">
        <img src="/images/star.jpg" style="z-index:-100; width:100%; height: 100%; opacity:0.6; position:absolute ">
        <div style="margin-left: 620px; font-size: 30px;">无源算力交易平台</div>
        <div style="border: solid #20B2AA 2px; margin-left: 9.72%;margin-right: 9.72%;"></div>
        <br>
        <table class="marketTable">
            <thead>
                <tr >
                    <th class="th"><NavTab :imgUrl="indexUrl[0][0]" :tableName="indexUrl[0][1]"></NavTab></th>
                    <th class="th"><NavTab :imgUrl="indexUrl[1][0]" :tableName="indexUrl[1][1]"></NavTab></th>
                    <th class="th"><NavTab :imgUrl="indexUrl[2][0]" :tableName="indexUrl[2][1]"></NavTab></th>
                    <th class="th"><NavTab :imgUrl="indexUrl[3][0]" :tableName="indexUrl[3][1]"></NavTab></th>
                    <th class="th"><NavTab :imgUrl="indexUrl[4][0]" :tableName="indexUrl[4][1]"></NavTab></th>
                    <th class="th"><NavTab :imgUrl="indexUrl[5][0]" :tableName="indexUrl[5][1]"></NavTab> </th>
                    <!-- <th class="th1"></th> -->
                </tr>
            </thead>
            <br>
            <!-- <tbody>
                <tr v-for="model in sellingNodes" :key="model.name">
                    <td class="td">live</td>
                    <td class="td">{{ model.name }}</td>
                    <td class="td">{{ model.model }}</td>
                    <td class="td">{{ model.version }}</td>
                    <td class="td">{{ model.price }}</td>
                    <td class="td"><button @click="buy(model.id)">+</button> {{ model.token*1000 }}<button @click="drop(model.id)">-</button> </td>
                    <td ><span v-html="'&nbsp;&nbsp;'"></span><button style="background-color:chartreuse;border-width:10px;border-color: chartreuse; border-radius: 10px;" @click="goToPay">确定购买</button></td>
                </tr>
            </tbody> -->
        </table>
        <!-- <video class="video" playsinline="" autoplay="" muted="" loop="">
                <source src="/images/市场背景.mp4" type="video/mp4">
        </video> -->
        <div v-for="item in sellingNodes" style="border:#FF8C00 5px solid;margin-left:340px; border-radius: 20px;margin-right: 340px;z-index: 20;">
            <div style ="position: relative;line-height:94px; display:inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#00FFFF;top: 5.7px;z-index:-2;">
                <img src="images/statuslive.png" style="width:120px;height: 70px;position: relative;"/>
                <div style="border: 1px solid #000; position: relative; margin-top: -37px;"></div>
                <div style="background: #F5F5F5;width:140px;height: 17.9px; font-size: 9px;line-height: 13px;">状态：<span style="font-size: 9px;font:900">活跃</span></div>
            </div>
            <div style ="position: relative;line-height:94px; display:inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#FF0000;top: 5.7px;z-index:-2;">
                <img :src="getImg(item.id)" style="width:120px;height: 70px;position: relative;"/>
                <div style="border: 1px solid #000; position: relative; margin-top: -37px;"></div>
                <div style="background: #808080;width:140px;height: 17.9px; font-size: 9px;line-height: 13px;">状态：<span style="font-size: 9px;font:900">活跃</span></div>
            </div>
            <!-- <div style ="position: relative;line-height:94px; display:inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#FF4500;top: -27.3px;">
                <img src="images/market/head1.png" style="width: 120px;height: 70px;"/>
                <div style="border: 1px solid #000; position: relative; margin-top: -37px;"></div>
            </div> -->
            <div style="position: relative;line-height:94px; display:inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#FFC0CB;top: -27.3px;">{{ item.model }}</div>
            <div style="position:relative;line-height:94px; display: inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#9370DB;top:-27.3px;">{{ item.version }}</div>
            <div style="position:relative;line-height: 94px; display: inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#FFFF00;top:-27.3px;">{{ item.price }}</div>
            <div style="position: relative;line-height: 94px; display: inline-block;font-weight: 1000; width:140px;height: 94px;text-align: center;background:#F5F5F5;top:-27.3px">
                <button class="button_s" @click="buy(item.id)">+</button> {{ item.token*1000 }}<button class="button_s" @click="drop(item.id)"> -</button>
            </div>
            <!-- <div style="position:relative;top:-46px; display: inline-block; text-align: center;background:#20B2AA;width: 140px;font-size: 12px;">状态:<span style="font-size:12px;font-weight: 700;">活跃</span></div>
            <div style="position:relative; top:-46px;display:inline-block; text-align: center;background:#FF4500;width: 140px;font-size: 12px;">节点:{{ item.name }}</div> -->
        </div>
        <br>
        <buttom @click="goToPay" @mouseover="changeBack" @mouseleave="hhh" :class="buttonStyle" style="line-height:40 px; height: 30px;width:80px; border: 5px #7FFFAA outset;border-radius: 12px;margin-left: 48%;">确认购买</buttom>
        <div>
            <video src="images/川流不息.mp4" class="video" ref="video" muted autoplay loop></video>
        </div>
    </div>
    
</template>

<script setup>
    import axios from 'axios'
    import NavTab from "@/components/NavTab.vue"
    import {ref,reactive,onMounted,set} from 'vue';
import { router } from '@/router';
    let sellingNodes = reactive([{ name: "admin", model: "Gemini-Pro", version: "1.5", price: 1, token: 100, id: 1 }])
    const buttonStyle = ref("button_h");
    
    const indexUrl = ref([
        ["images/market/status.png","状态"],
        ["images/market/user.png","节点名称"],
        ["images/market/category.png","模型种类"],
        ["images/market/version.png","模型版本"],
        ["images/market/money.png","模型收费(千字)"],
        ["images/购买.png","购买数量(千字)"],
    ])
    let modelId = 1
    function getImg(id) {
        console.log()
        return 'images/market/head'+id+'.png'
    }
    function buy(id) {

        // if (id==0) {
        //     return
        // }

            sellingNodes[id].token++;
    
        
    }
    function drop(id) {
        if (sellingNodes[id].token>0){
            sellingNodes[id].token--
        } else {
            alert("不能再少了哟！！！")
        }
    }
    function init() {
        // console.log(sellingNodes.value)
        axios.get("http://localhost:5678/api/map/search").then(
            res => {
                console.log(res.data.loc.length)
                let i = 1;
                for ( ;i<res.data.loc.length;i++){
                    let temp = res.data.loc[i]
                    console.log(temp)
                    // sellingNodes.push({name:temp.name,model:temp.api.platform,version:temp.api.version,price:temp.price,token:0,id:modelId++})
                    // sellingNodes.push({ name: temp.name, model: temp.api.platform, version: temp.api.version, price: temp.price, token: 0, id: modelId++ })
                    let model = { name: "admin", model: "Gemini-Pro", version: "1.5", price: 0, token: 0, id: modelId }
                    modelId=modelId+1
                    model.name = temp.name
                    model.model = temp.api.platform
                    model.version = temp.api.version
                    model.id=modelId
                    model.price=temp.price
                    model.token
                    sellingNodes.push(model)
                    console.log(sellingNodes)
                    // set(sellingNodes, i, { name: temp.name, model: temp.api.platform, version: temp.api.version, price: temp.price, token: 0, id: modelId++ });
                }
            }
        )
    }
    function goToPay() {
        router.push("/payment")
    }
    init()
    function changeBack() {
        buttonStyle.value = "button_c"
    }
    function hhh() {
        buttonStyle.value = "button_h"
    }

</script>

<style scoped>

/* .button{
    background:#00CED1;
    font-weight: 1000;
    border: 3px solid #000000;
    border-radius: 5px;
    height: ;
} */
.video{
    content: "";
    position: absolute;
    width:100%;
    height: 100%;

    display:block;
    z-index:-10;
    top:0;
    left:0;
    background:#000000;
    opacity: 0.7;
}

.marketList{
    width:100%;
    height:100%;
    top:0;
    left:0;
    overflow:hidden;
    position: absolute;
    z-index:-50

}

.marketTable{
    /* border:20px solid #e9e9e9;  */
    border-collapse:collapse;
    border-spacing: 0;
    margin:0 auto;
}

.th{
    padding: 0;
    border: 0.5px solid #000000;
    text-align:center;
    font-weight:1100;
    overflow:hidden;
    width: 140px;
    
}
.th1{
    padding: 0;
    border: 0.5px ;
    text-align:center;
    font-weight:700;
    border-radius: 10px;
    overflow:hidden;
    width: 50px;
}
.td{
    /* padding: 8px 16px; */
    /* border: 1px solid #00BFFF; */
    text-align:center;
    background:#ffffff;
}
.button_s{
    border:0.4px solid #000;
    background: #00FFFF;
    height: 20px;
    width: 20px;
    position: relative;
    margin-top:px;
    z-index: 10;
    line-height: 10px;
    text-align: center;
}
.button_h{
    display: block;
    text-align: center;
    background: #00FA9A;
}
.button_c{
    display: block;
    text-align: center;
    background: #FF0000;
}

</style>