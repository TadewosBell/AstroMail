<script setup>
import { reactive } from 'vue';
import { Launch_Smtp_Server } from '../../wailsjs/go/main/App'

const emit = defineEmits(['NextSlide'])
function NextSlide(folder) {
  emit('NextSlide', folder)
}

const data = reactive({
    Username: 'tadewosbell',
    Domain: 'astrocommits.com', // null,
    AwsID: 'AKIA3A3TZCPGSWECKY7I', // null,
    AwsSecret: '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', // null,
})

// Function to validate the domain
const isValidDomain = (domain) => {
    // Simple regex for domain validation - this can be adjusted as needed
    const pattern = /^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9]\.[a-zA-Z]{2,}$/;
    console.log(pattern.test(domain))
    return pattern.test(domain);
};

// Launch function with validations
const Launch = () => {
    const Next_Slide = NextSlide;
    // Check if Domain, AwsID, and AwsSecret are not empty and if Domain is valid
    console.log(data.AwsID, data);
    if (data.Domain && data.AwsID && data.AwsSecret && isValidDomain(data.Domain)) {
    Launch_Smtp_Server(data.Username, data.Domain, data.AwsID, data.AwsSecret).then(result => {
    }).catch(error => {
        console.error('Launch failed:', error);
        // Handle the error appropriately
    });
    } else {
    // If validation fails, log an error or handle it as necessary
    console.error('Validation failed: Make sure all fields are filled correctly.');
    return;
    }
    NextSlide();
};

</script>
<template>
        <div class="div1">
            <img class="slide_image" alt="Wails logo" src="../assets/images/setup3.png" />
        </div>
        <div class="div2">
            <h1>
              Step 3: Launch!
            </h1>
            <p>Enter your domain, AWS ID and Secret Key and click deploy</p>
            <input v-model="data.Username" class="setupInput" type="text" placeholder="email username" >
            <br/>
            <input v-model="data.Domain" class="setupInput" type="text" placeholder="Domain" >
            <br/>
            <input v-model="data.AwsID" class="setupInput" type="text" placeholder="AWS ID" >
            <br />
            <input v-model="data.AwsSecret" class="setupInput" type="text" placeholder="AWS Secret Key" >
            <br />
            <button class="next" v-on:click="Launch">Launch</button>
        </div>
</template>
<style>
.slide_image {
    width: 100%;
}

.next {
    position: absolute;
    font-size: 15px;
    font-family: Arial;
    width: 140px;
    height: 50px;
    border-width: 1px;
    color: #fff;
    border-color: #0b0e07;
    border-top-left-radius: 3px;
    border-top-right-radius: 3px;
    border-bottom-left-radius: 3px;
    border-bottom-right-radius: 3px;
    box-shadow: inset 0px -3px 7px 0px #29bbff;
    text-shadow: inset 0px 1px 0px #263666;
    background: linear-gradient(#2dabf9, #0688fa);
    bottom: 0;
    right: 10px;
}

.setupInput {
    width: 50%;
    height: 25px;
    margin-top: 10px;
}

.div2 {
    background-color: #f2f6f6;
    width: 100%;
    position: relative;
}
</style>