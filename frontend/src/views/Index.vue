<template>
    <!-- <div class="index-page" :style="{height: winHeight+'px'}"> -->
    <div class="index-page" :style="{ backgroundImage: `url(${backgroundImage})`,height: winHeight+'px','background-size':'cover'}" >
      
      <div class="menu-box">
        <el-menu
            mode="horizontal"
            :ellipsis="false"
        >
          <div class="menu-item">
            <el-image :src="logo" alt="Geek-AI"/>
            <div class="title">{{ title }}</div>
          </div>
          <div class="menu-item">
            <el-image :src="voiceImages" @click="changeVoice" class="voicePic"/>
            <div class="title">音乐</div>
          </div>
        </el-menu>
      </div>
      <div class="content">
        <h1 style="margin-left: 85px; color:#ffffff;"> {{ title }}</h1>
        <p>{{ slogan }}</p>
        <div class="button-group">
      <el-button @click="router.push('/chat')" type="primary" icon="chat">聊天界面</el-button>
      <el-button @click="router.push('/map')" type="success" icon="mj">模型全网分布</el-button>
      <el-button @click="router.push('/market')" type="info" icon="sd">大语言模型购买</el-button>
      <el-button @click="router.push('/xmind')" type="warning" icon="xmind">敬请期待</el-button>
      </div>
        <div id="animation-container"></div>
      </div>
  
      <div class="footer">
        <footer-bar />
      </div>
    </div>
  </template>
  
  

  <script setup>
  import * as THREE from 'three';
 
  import {onMounted, ref} from "vue";
  import {useRouter} from "vue-router";
  import FooterBar from "@/components/FooterBar.vue";
  import {httpGet} from "@/utils/http";
  import { getLocalIP } from '@/utils/worker';
  import {ElMessage} from "element-plus";
  import Howler from 'howler';
  import { onBeforeRouteLeave } from 'vue-router';
  
  const musicUrl = ref('music/index.mp3');
  const sound = ref(new Howler.Howl({ src: [musicUrl.value], loop: true }));
  const backgroundImage = ref('images/背景1.png');
  
  const router = useRouter()
  
  const title = ref("数悦深者—区块链赋能的新时代无痕AI模型")
  const logo = ref("/images/logo2.png")
  const slogan = ref("'链接'你我的语言，'链接'世界的AI！")
  const size = Math.max(window.innerWidth * 0.5, window.innerHeight * 0.8)
  const winHeight = window.innerHeight-150
  const images = ['images/背景2.png','images/背景3.png','images/背景4.png','images/背景5.png'];
  let intervalId = null;
  const voiceImages = ref("images/open_voice.png");
  const voiceStatus = ref("true");
  
  
  onBeforeRouteLeave((to, from, next) => {
    sound.value.stop();
    next()
  });
  
  onMounted(() => {
    
    intervalId = setInterval(() => {
      const index = (images.indexOf(backgroundImage.value) + 1) % images.length;
      backgroundImage.value = images[index];
    }, 2000);
    sound.value.play();
    
  
    httpGet("/api/config/get?key=system").then(res => {
      title.value = res.data.title
    }).catch(e => {
      ElMessage.error("获取系统配置失败：" + e.message)
    })
    init()
  })
  
  // onUnmounted(() => {
  //   // sound.value.stop()
  // });
  
  
  const init = () => {
    // 创建场景
    // 创建场景
  
    const scene = new THREE.Scene();
  
    // 创建相机
    const camera = new THREE.PerspectiveCamera(30, 1, 0.1, 1000);
    camera.position.z = 3.88;
  
    // 创建渲染器
    const renderer = new THREE.WebGLRenderer({ antialias: true, alpha: true });
    renderer.setSize(size, size);
    renderer.setClearColor(0x000000, 0);
    const container = document.getElementById('animation-container');
    container.appendChild(renderer.domElement);
  
    // 加载地球纹理
    const loader = new THREE.TextureLoader();
    // 加载纹理图片
    loader.load(
    '/images/sky.png',
    function (texture) {
    // 创建地球球体
    const geometry = new THREE.SphereGeometry(1, 32, 32); // 创建球体的几何体，参数分别为半径、水平和垂直切片数
    const material = new THREE.MeshPhongMaterial({ // 创建球体的材质
    map: texture, // 使用纹理图片作为颜色贴图
    bumpMap: texture, // 使用同一张纹理作为凹凸贴图
    bumpScale: 0.05, // 调整凹凸贴图的影响程度
    specularMap: texture, // 高光贴图
    specular: new THREE.Color('#007bff'), // 高光颜色
    });
    const earth = new THREE.Mesh(geometry, material); // 创建球体对象
    scene.add(earth); // 将球体添加到场景中
    // 添加环境光和点光源
    const ambientLight = new THREE.AmbientLight(0xffffff, 0.3); // 创建环境光，颜色为白色，强度为0.3
    scene.add(ambientLight); // 将环境光添加到场景中
    const pointLight = new THREE.PointLight(0xffffff, 0.8); // 创建点光源，颜色为白色，强度为0.8
    pointLight.position.set(5, 5, 5); // 设置点光源的位置
    scene.add(pointLight); // 将点光源添加到场景中
  
    // 创建动画
    const animate = function () {
      requestAnimationFrame(animate); // 循环调用animate函数，实现动画效果
  
      // 使地球自转和公转
      earth.rotation.y += 0.01; // 地球绕Y轴旋转
  
      renderer.render(scene, camera); // 渲染场景和相机
  
    };
  
    // 执行动画
    animate(); // 调用animate函数开始执行动画
    }
  );
  
  }
  
  
  
  function changeVoice() {
      if (voiceStatus.value==true) {
        voiceStatus.value = false
        sound.value.stop()
        voiceImages.value = "images/close_voice.png";
      } else {
        voiceStatus.value = true
        sound.value.play()
        voiceImages.value = "images/open_voice.png";
      }
    }
  
  </script>
  
  <style scoped>
  .button-group {
    display: flex;
    justify-content: space-around;
    margin-top: 20px;
  }
  </style>
  
  <style lang="stylus" scoped>
  @import '@/assets/iconfont/iconfont.css'
  .index-page {
    margin: 0
    background-color #000000 /* 科技蓝色背景 */

    overflow hidden
    color #ffffff
    display flex
    justify-content center
    align-items baseline
    padding-top 150px
  
    .bg {
      position absolute
      top 0
      left 0
      width 100vw
      height 100vh
      background-image url("~@/assets/img/ai-bg.jpg")
      //filter: blur(8px);
      background-size: cover;
      background-position: center;
    }
  
    .menu-box {
      position absolute
      top 0
      width 100%
      display flex
  
      .el-menu {
        padding 0 30px
        width 100%
        display flex
        justify-content space-between
        background none
        border none
  
        .menu-item {
          display flex
          padding 20px 0
  
          color #ffffff
          .voicePic {
            width: 3.125rem
            height: 1.25rem
          }
  
          .title {
            font-size 24px
            padding 10px 10px 0 10px
          }
  
          .el-image {
            height 50px
          }
  
          .el-button {
            margin-left 10px
  
            span {
              margin-left 5px
            }
          }
        }
      }
    }
  
    .content {
      text-align: center;
      position relative
  
      h1 {
        font-size: 5rem;
        margin-bottom: 1rem;
      }
  
      p {
        font-size: 1.5rem;
        margin-bottom: 2rem;
      }
  
      .el-button {
        padding: 25px 20px;
        font-size: 1.3rem;
        transition: all 0.3s ease;
  
        .iconfont {
          font-size 1.6rem
          margin-right 10px
        }
      }
  
      #animation-container {
        display flex
        justify-content center
        width 100%
        height: 300px;
        position: absolute;
        top: 350px
  
      }
    }
  
    .footer {
      .el-link__inner {
        color #ffffff
      }
    }
  
  }
  </style>