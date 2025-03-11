<script setup>
import "bootstrap-icons/font/bootstrap-icons.css";
</script>

<template>
  <div>
    <h2><i class="bi bi-info-circle text-success"></i>
         Instructions</h2>
    <p class="text-body-secondary">This page has tools for working with <a href="https://github.com/sjpiper145/MakerSkillTree" target="_blank">Maker Skill Trees</a>.</p>

    <ol>
        <li>Select a tool from the list above.</li>
        <li>Upload skill tree files (SVG or YAML) that you created with the 
            <a href="https://schme16.github.io/MakerSkillTree-Generator/" target="_blank">Maker Skill Tree Generator</a> 
            or downloaded from the <a href="https://github.com/sjpiper145/MakerSkillTree" target="_blank">Maker Skill Tree repository</a>.
        </li>
        <li>View or download the results.</li>
    </ol>
    <br />
    <h4>Can I make a suggestion?</h4>

    <p>Please do! <a href="https://github.com/josephlewis42/skilltreetool/issues/new">You can open an issue on GitHub with suggestions.</a></p>

    <br />
    <h4>Is my data private?</h4>

    <p>Yes! All processing happens in your browser using WebAssembly.</p>

    <p>There's also <a href="https://github.com/josephlewis42/skilltreetool" target="_blank">a CLI version of the tool with more features</a>.</p>
  </div>
</template>

<style scoped>
</style>
