<script setup>
import FileUpload from './FileUpload.vue';
import { useTemplateRef, ref, computed, defineProps, defineEmits } from 'vue'


var beforeText = ref(null);
var afterText = ref(null);


var output = computed(() => {
  if (!beforeText.value || !afterText.value) {
    return "Diff will show up here once you select files above...";
  }

  const result = diff(beforeText.value, afterText.value);

  const decoded = JSON.parse(result)

  if(decoded.err) {
    return decoded.err
  }
  

  return decoded.data
});

var beforeWarning = ref(null);
var afterWarning = ref(null);

function updateBefore(text) {
  beforeText.value = text
}


function updateAfter(text) {
  afterText.value = text
}

</script>

<template>
  <div>

    <div class="container">
      <div class="row">
        <h2><i class="bi bi-file-earmark-diff"></i> Diff</h2>
        <p class="text-body-secondary">Shows the changes between two files in a <a href="https://keepachangelog.com/en/1.0.0/" target="_blank">Keep a Changelog</a> style format.</p>

      </div>
      <div class="row">
        <div class="col">
          <p>
            <b>Before: </b>
          <span class="text-body-secondary">Choose the original file you want to compare changes against.</span>
          </p>
        </div>
        <div class="col">
          <p><b>After: </b>
            <span class="text-body-secondary"> Choose the new file that has the changes.</span>
          </p>
        </div>
      </div>
      <div class="row">
        <div class="col">
          <FileUpload @fileDataLoaded="updateBefore" accept=".svg,.yaml,.yml"></FileUpload>
          <div class="text-danger" v-if="beforeWarning">{{ beforeWarning }}</div>
        </div>
        <div class="col">
          <FileUpload @fileDataLoaded="updateAfter" accept=".svg,.yaml,.yml"></FileUpload>
          <div class="text-danger" v-if="afterWarning">{{ afterWarning }}</div>
        </div>
      </div>
      <div class="row pt-4">
        <div class="col">
          <label for="diffOutput"><b>Diff results</b></label>
          <textarea id="diffOutput" class="form-control" placeholder="Diff will show up here once you select files above..." :value="output" style="width:100%; height:50vh;" readonly></textarea>

        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
</style>
