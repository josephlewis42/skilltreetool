<script setup>
import { uniqueId } from '@/assets/uid';
import { useTemplateRef, ref, computed, defineProps, defineEmits } from 'vue'
const emit = defineEmits(["filePicked", "fileDataLoaded"]);
const props = defineProps(["accept"])

const isDragging = ref(false);
const fileUpload = useTemplateRef("fileUpload");
const file = ref(null);
const loadErr = ref(null);
const isLoading = ref(false);
const uid = uniqueId();

function onUpload() {
    loadFile(fileUpload.value.files ? fileUpload.value.files[0] : null)
}

function dragover(e) {
    isDragging.value = true;
}

function dragleave() {
    isDragging.value = false;
}

function drop(e) {
    loadFile(e.dataTransfer.files[0])
}

function loadFile(loadedFile) {
    isDragging.value = false
    loadErr.value = null

    if (loadedFile == null) {
        file.value = null
        fileUpload.value.value = null
        emit("fileDataLoaded", null, null)
        return
    }

    file.value = loadedFile

    // TODO: show loading
    isLoading.value = true

    loadedFile.text()
        .then((text) => {
            emit("fileDataLoaded", text, loadedFile.name)
        })
        .catch((e) => {
            loadErr.value = e
            emit("fileDataLoaded", null, null)
        })
        .finally(() => {
            isLoading.value = false
        });
}

</script>

<template>
  <div class="card" :class="{'border-2 border-info': isDragging, 'border-success': file}" @dragover.prevent="dragover" @dragleave.prevent="dragleave" @drop.prevent="drop" @mouseleave="dragleave">
    <div class="card-body">
        <div >
            <input style="display:none;" :id="uid" type="file" name="file" @change="onUpload" ref="fileUpload" :accept="accept" />
            <label :for="uid">
                <div v-show="isLoading">Loading...</div>
                <div v-if="file">File uploaded: <code>{{ file.name }}</code> <i class="bi bi-x-circle" @click.prevent="loadFile(null)"></i></div>
                <div v-if="isDragging">Drop to upload</div>
                <div v-else>Drop files here or <u>click to upload</u>.</div>
            </label>
        </div>

        <div class="text-danger" v-if="loadErr">Error loading data: {{ loadErr }}</div>
    </div>
  </div>
</template>

<style scoped>
</style>
