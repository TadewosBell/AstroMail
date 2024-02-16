<script setup>
import { reactive, onMounted } from 'vue'
import sidenavHeader from '../components/sidenav-header.vue';
import topHeader from '../components/header.vue';
import itemList from '../components/item-list.vue';
import sidenav from '../components/sidenav.vue';
import content from '../components/content.vue';
import { Get_Inbox, Get_Sent, Get_Next_Page } from '../../wailsjs/go/main/App'
import { EventsOn } from '../../wailsjs/runtime/runtime';


EventsOn('emailSent', () => {
  GetSent();
});

const data = reactive({
  folder: 'inbox',
  folders: {
    'inbox': [],
    'sent': [],
  },
  focused_item: null,
  currentPage: 1,
})

const emit = defineEmits(['composeEmail'])

function GetInbox() {
  data.folders.inbox = [];
  Get_Inbox().then(result => {
    console.log(result);
    let currentId = 0;
    result.forEach((email) => {
      // Check if email string is not empty before parsing
      if (email.trim() !== "") {
        try {
          const emailJson = JSON.parse(email);
          console.log(emailJson);
          if (emailJson.from) {
            emailJson.id = currentId++;
            data.folders.inbox.push(emailJson);
          }
        } catch (e) {
          console.error("Failed to parse email JSON:", e);
          // Handle the error or ignore the email
        }
      }
    });
  }).catch(error => {
    console.error("Error fetching inbox:", error);
  });

}

function GetSent() {
  let folders = {
    inbox: data.folders.inbox,
  }
  Get_Sent().then(result => {
    let sent = [];
    let currentId = 0;
    result.forEach((email) => {
      const emailJson = JSON.parse(email);
      if (typeof emailJson === "object") {
        emailJson.id = currentId;
        sent.push(emailJson);
        currentId = currentId + 1;
      }
    });
    // Move the assignment inside the then block
    folders.sent = sent;
    data.folders = {
    ...folders,
    }
  }).catch(error => {
    console.error("Error fetching sent emails:", error);
  });
}

onMounted(() => {
  GetInbox()
  GetSent()
})


function ItemSelected(index) {
  data.focused_item = index;
}

function ChangeFolder(folder) {
  console.log(folder);
  data.folder = folder;
  data.focused_item = null;
}

function ComposeEmail() {
  emit('composeEmail')
}

function PreviousPage(){

}

function NextPage() {
  let nextPage = data.currentPage + 1;
  Get_Next_Page(data.folder, nextPage).then(result => {
    console.log(result)
    if (result.length >0){
      data.folders[data.folder] = [];
      data.currentPage = nextPage;
      result.forEach((email) => {
      // Check if email string is not empty before parsing
      if (email.trim() !== "") {
        try {
          const emailJson = JSON.parse(email);
          console.log(emailJson);
          if (emailJson.from) {
            emailJson.id = currentId++;
            data.folders[data.folder].push(emailJson);
          }
        } catch (e) {
          console.error("Failed to parse email JSON:", e);
          // Handle the error or ignore the email
        }
      }
    });
    }
  }).catch(error => {
    console.error("Error fetching sent emails:", error);
  });
}

</script>

<template>
  <main class="parent">
    <sidenav-header />
    <top-header @previous-page="PreviousPage" @next-page="NextPage" @refresh-email="GetInbox" @compose-email="ComposeEmail" :title="data.folder" :currentPage="data.currentPage" />
    <item-list @item-selected="ItemSelected" :folder="data?.folders[data.folder]" :folderName="data.folder" />
    <sidenav @folder-selected="ChangeFolder" />
    <content :email="data?.folders[data.folder][data.focused_item? data.focused_item: 0]" />
  </main>
</template>

<style scoped>
.search {
  padding: 12px 16px 12px 16px;
  margin: 0 0 0 20px;
}


.parent {
  display: grid;
  grid-template-columns: repeat(2, 1.5fr) 0.1fr repeat(4, 1fr) 0.1fr 0fr repeat(3, 2.25fr);
  grid-template-rows: 1fr 0.1fr repeat(10, 1fr);
  grid-column-gap: 0px;
  grid-row-gap: 0px;
  height: 100vh;
  widows: 100%;
}

.div1 {
  grid-area: 1 / 1 / 2 / 3;
  overflow: hidden;
  background-color: #dfe3e3;
}

.div2 {
  grid-area: 1 / 3 / 2 / 13;
  width: 100%;
  display: flex;
  align-items: center;
  padding: 0 20px;
  background-color: #f2f6f6;
  --wails-draggable: drag;
  border-bottom: 1px solid rgba(5, 5, 5, 0.06);
}

.div3 {
  grid-area: 2 / 3 / 13 / 9;
  /* background-color: blue; */
  border-right: 1px solid rgba(5, 5, 5, 0.06);
  background-color: #fff;
}

.div4 {
  grid-area: 2 / 1 / 13 / 3;
  overflow: hidden;
  background-color: #dfe3e3;
}

.div5 {
  grid-area: 2 / 10 / 13 / 13;
  background-color: #fff;
}


.items {
  margin: 0;
  padding: 0;
  color: black
}

/* Style the list items */
.email {
  cursor: pointer;
  position: relative;
  padding: 12px 8px 12px 40px;
  list-style-type: none;
  background: rgb(255, 255, 255);
  font-size: 18px;
  transition: 0.2s;

  /* make the list items unselectable */
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  border-bottom: 2px solid rgba(5, 5, 5, 0.06);
}

.email:hover {
  background: #dfe3e3;
}

.folders {
  cursor: pointer;
  position: relative;
  padding: 05px 0px 5px 0px;
  list-style-type: none;
}
</style>
