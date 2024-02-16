<script setup>

import { BrowserOpenURL } from '../../wailsjs/runtime'

const emit = defineEmits(['NextSlide'])
function NextSlide(folder) {
    emit('NextSlide', folder)
}

const OpenAws = () => {
    console.log("Here")
    BrowserOpenURL('http://aws.amazon.com')
}
</script>
<template>
    <div class="div1">
        <img class="slide_image" alt="Wails logo" src="../assets/images/setup1.png" />
    </div>
    <div class="div2">
        <div class="astro-mail-setup">
            <h1>Welcome to AstroMail Email Server Setup</h1>
            <div class="step">
                <h3>Step 1: Create an AWS Account</h3>
                <ol>
                    <li>Sign up for AWS by visiting <button class="aws_button" @click="OpenAws">Amazon web services</button>
                    </li>
                    <li>
                        Click on "Sign In to the Console" and follow the instructions to create your account. You'll need to
                        provide some basic information and billing details.</li>
                    <li>Click Next when you have created an account</li>
                    </ol>
            </div>

            <button class="next" v-on:click="NextSlide">Next</button>
        </div>
    </div>
</template>
<style>
.slide_image {
    width: 100%;
}

.aws_button {
    font-size: 11px;
    font-family: Arial;
    width: 150px;
    height: 30px;
    border-width: 1px;
    color: rgba(0, 0, 0, 1);
    border-color: rgba(255, 153, 0, 1);
    font-weight: bold;
    border-top-left-radius: 28px;
    border-top-right-radius: 28px;
    border-bottom-left-radius: 28px;
    border-bottom-right-radius: 28px;
    text-shadow: 1px 1px 0px rgba(255, 153, 0, 1);
    background: rgba(255, 153, 0, 1);
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

.div2 {
    background-color: #f2f6f6;
    width: 100%;
    position:relative;
}

.perm {
    padding-left: 150px
}
</style>