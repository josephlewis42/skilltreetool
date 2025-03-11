<script setup>
import Instructions from './components/Instructions.vue'
import YamlToSvg from './components/YamlToSvg.vue'
import SvgToYaml from './components/SvgToYaml.vue'
import Diff from './components/Diff.vue'
import { ref, computed } from 'vue'
import "./assets/golang.js"

const page = ref("home")

function changePage(name) {
  page.value = name;
}

</script>

<template>
  <div>
    <header class="d-flex justify-content-center py-3">
      <ul class="nav nav-pills">
         
        <li class="nav-item"><a href="#" class="nav-link" :class="{ active: page == 'home' }"
            @click="changePage('home')"><i class="bi bi-info-circle"></i> Instructions</a></li>
        <li class="nav-item"><a href="#" class="nav-link" :class="{ active: page == 'diff' }"
            @click="changePage('diff')"><i class="bi bi-file-earmark-diff"></i> Diff</a></li>
        <li class="nav-item"><a href="#" class="nav-link" :class="{ active: page == 'yaml2svg' }"
            @click="changePage('yaml2svg')"><i class="bi bi-filetype-svg"></i> YAML to SVG</a></li>
        <li class="nav-item"><a href="#" class="nav-link" :class="{ active: page == 'svg2yaml' }"
            @click="changePage('svg2yaml')"><i class="bi bi-filetype-yml"></i> SVG to YAML</a></li>
      </ul>
    </header>

    <Instructions v-show="page == 'home'"></Instructions>
    <YamlToSvg v-show="page == 'yaml2svg'"></YamlToSvg>
    <SvgToYaml v-show="page == 'svg2yaml'"></SvgToYaml>
    <Diff v-show="page == 'diff'"></Diff>
  </div>

</template>

<style scoped></style>
