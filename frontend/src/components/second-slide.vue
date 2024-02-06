<script setup>
import VueJsonPretty from 'vue-json-pretty';
import 'vue-json-pretty/lib/styles.css';
const emit = defineEmits(['NextSlide'])
function NextSlide(folder) {
  emit('NextSlide', folder)
}


const permJson = {
                "Version": "2012-10-17",
                "Statement": [
                    {
                        "Effect": "Allow",
                        "Action": [
                            "ses:VerifyDomainIdentity",
                            "ses:GetIdentityVerificationAttributes",
                            "ses:SetIdentityMailFromDomain"
                        ],
                        "Resource": "*"
                    }
                ]
            }
</script>
<template>
        <div class="div1">
            <img class="slide_image" alt="Wails logo" src="../assets/images/setup2.png" />
        </div>
        <div class="div2">
      <h2>Step 2: Create an IAM User</h2>
      <ol>
        <li>Log into the <button class="aws_button" @click="OpenAws">Amazon web services</button></li>
        <li>Find and select <strong>IAM (Identity and Access Management)</strong> from the Services menu.</li>
        <li>Follow the instructions to create a new IAM user.</li>
        <li>When setting permissions, use the provided JSON to attach the necessary permissions directly:</li>
      </ol>
      <div class="perm">
            <vue-json-pretty :data="permJson" />
        </div>
            <button class="next" v-on:click="NextSlide" >Next</button>
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

.div2 {
    background-color: #f2f6f6;
    width: 100%;
    position: relative;
}
</style>