<script setup>
import FileUpload from './FileUpload.vue';
import { useTemplateRef, ref } from 'vue'

const output = ref(null)
const err = ref(null)
const preview = useTemplateRef("preview")
const downloadLink = useTemplateRef("download")

function uploadFile(text, name) {
  output.value = null
  err.value = null
  downloadLink.value.download = ""
  downloadLink.value.href = ""

  if (!text) {
    return
  }

  const decoded = JSON.parse(yaml2svg(text))

  if (decoded.err) {
    err.value = decoded.err
    return
  } else {
    output.value = decoded.data

    var blob = new Blob([decoded.data], {type: "text/html; charset=utf-8"});
    const objectURL = URL.createObjectURL(blob)
    preview.value.src = objectURL
    downloadLink.value.download = name + ".svg";
    downloadLink.value.href = objectURL
  }
}
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h2><i class="bi bi-filetype-svg"></i> YAML to SVG</h2>
        <p class="text-body-secondary">Convert 
          <a href="https://github.com/sjpiper145/MakerSkillTree" 
            target="_blank">Maker Skill Tree</a> 
          YAML to 
          <a href="https://schme16.github.io/MakerSkillTree-Generator/"
            target="_blank">Maker Skill Tree Generator</a> compatible SVGs.</p>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">Input</h5>
            <label>Select the YAML to load</label>

            <FileUpload @fileDataLoaded="uploadFile" accept=".yml,.yaml"></FileUpload>
          </div>
        </div>

      </div>

      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title d-flex justify-content-between align-items-center">
              SVG Output
              <a
                class="btn btn-sm btn-primary" :class="{'btn-primary': output, 'btn-secondary': !output}" ref="download">Download</a>
            </h5>
            <form>
              <div class="form-group">
                <div class="alert alert-warning" role="alert" v-if="err">
                  {{ err }}
                </div>

                <div v-show="output == null">
                  <p class="lead text-center">
                    Output will show up here once you load a file.
                  </p>
                </div>

                <div v-show="output">
                  <p class="card-text">Preview</p>
                  <div class="card border-dark">
                    <iframe ref="preview" class="" style="width:100%; height:50vh;" sandbox></iframe>
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>

      </div>
    </div>
  </div>

</template>

<style scoped></style>
