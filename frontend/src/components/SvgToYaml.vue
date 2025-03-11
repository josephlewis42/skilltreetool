<script setup>
import FileUpload from './FileUpload.vue';
import { useTemplateRef, ref } from 'vue'

var output = ref(null)
var err = ref(null)
const downloadLink = useTemplateRef("download")

function uploadFile(text, name) {
  output.value = null
  err.value = null
  downloadLink.value.download = ""
  downloadLink.value.href = ""


  if(! text) {
    return
  }

  const decoded = JSON.parse(svg2yaml(text))

  if(decoded.err) {
    err.value = decoded.err
    return
  } else {
    output.value = decoded.data
    var blob = new Blob([decoded.data], {type: "text/html; charset=utf-8"});
    const objectURL = URL.createObjectURL(blob)
    downloadLink.value.download = name + ".yaml";
    downloadLink.value.href = objectURL

  }
}

</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col">
        <h2><i class="bi bi-filetype-yml"></i> SVG to YAML</h2>
        <p class="text-body-secondary">Convert <a href="https://schme16.github.io/MakerSkillTree-Generator/"
            target="_blank">Maker Skill Tree Generator</a> compatible SVGs to <a
            href="https://github.com/sjpiper145/MakerSkillTree" target="_blank">Maker Skill Tree</a> YAML.</p>
      </div>
    </div>
    <div class="row">
      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title">Input</h5>

            <label>Select the SVG to load</label>

            <FileUpload @fileDataLoaded="uploadFile" accept=".svg"></FileUpload>
          </div>
        </div>

      </div>

      <div class="col">
        <div class="card">
          <div class="card-body">
            <h5 class="card-title d-flex justify-content-between align-items-center">
              YAML Output
              <a
                class="btn btn-sm btn-primary" :class="{'btn-primary': output, 'btn-secondary': !output}" ref="download">Download</a>
            </h5>
            <form>
              <div class="form-group">
                <div class="alert alert-warning" role="alert" v-if="err">
                  {{ err }}
                </div>

                <div v-show="output != null">
                  <textarea class="form-control" placeholder="Waiting for file..." :value="output" style="width:100%; height:50vh;" readonly></textarea>
                </div>
                <div v-show="output == null">
                  <p class="lead text-center">
                    Output will show up here once you load a file.
                  </p>
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
