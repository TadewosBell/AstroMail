<script setup>
import { reactive, onMounted } from 'vue'
import sidenavHeader from '../components/sidenav-header.vue';
import topHeader from '../components/header.vue';
import itemList from '../components/item-list.vue';
import sidenav from '../components/sidenav.vue';
import content from '../components/content.vue';
import { parseDateToJson } from "../../util/time"

import { Refresh_Inbox, Get_Items } from '../../wailsjs/go/main/App'
import { EventsOn } from '../../wailsjs/runtime/runtime';


EventsOn('Sent', () => {
  console.log("ALREADY SENT")
  
});

EventsOn('Sending', () => {
  console.log("ALREADY SENDING")
});

EventsOn('SendFail', () => {
  console.log("Sent failed")
});


let data = reactive({
  folder: 'inbox',
  folders: {
    'inbox': [],
    'sent': [],
  },
  focused_item: 0,
  current_page: 1,
  current_item: null,
})

const emit = defineEmits(['composeEmail'])

function sortByParsedDate(emails) {
  emails.sort((a, b) => {
    const dateA = a.parsedDate ? new Date(a.parsedDate.dateTime) : null;
    const dateB = b.parsedDate ? new Date(b.parsedDate.dateTime) : null;

    // Convert dates to timestamps for comparison
    const timestampA = dateA ? dateA.getTime() : 0;
    const timestampB = dateB ? dateB.getTime() : 0;

    // Handle cases where parsedDate or parsedDate.dateTime is undefined
    if (timestampA && timestampB) {
      return timestampB - timestampA;
    } else if (!timestampA && timestampB) {
      return 1; // Put b before a if a has no parsedDate
    } else if (timestampA && !timestampB) {
      return -1; // Put a before b if b has no parsedDate
    } else {
      return 0; // Leave the order unchanged if both have no parsedDate
    }
  });
  emails.forEach((item, index) => {
    item.id = index;
  })
}


function GetItems(folder, page) {
  let itemsList = [];
  Get_Items(folder, page).then(result => {
    console.log(result)
    if (result.length > 0) {
      data.current_page = page;
      let currentId = 0;
      result.forEach((email) => {
        // Check if email string is not empty before parsing
        if (email.trim() !== "") {

          try {
            const emailJson = JSON.parse(email);
            emailJson['parsedDate'] = parseDateToJson(emailJson.date);
            if (emailJson.from) {
              emailJson.id = currentId++;
              itemsList.push(emailJson)
              // data.folders[folder].push(emailJson);
            }
          } catch (e) {
            console.error("Failed to parse email JSON:", e);
            // Handle the error or ignore the email
          }
        }
      });
    }
    sortByParsedDate(itemsList)
    data.folders = {
      ...data.folders,
      [folder]: itemsList
    }
    // data.folders[folder] = sortByParsedDate(data.folders[folder]);
  }).catch(error => {
    console.error("Error fetching items:", error);
  });
}

onMounted(() => {
  refreshItems();
  ItemSelected(0);
})

function refreshItems() {
  Refresh_Inbox().then(result => {
    GetItems('inbox', 1);
  }).catch(error => {
    console.error("Error fetching items:", error);
  });
}

function ItemSelected(index) {
  
  data.focused_item = index;
  const current_item = data?.folders[data.folder][index];
  data.current_item = current_item;
}

function ChangeFolder(folder) {
  data.folder = folder;
  data.focused_item = 0;
  data.current_item = data?.folders[data.folder][0];
  GetItems(folder, 1);
}

function ComposeEmail() {
  emit('composeEmail')
}

function PreviousPage() {
  let prevPage = data.current_page - 1;
  GetItems(data.folder, prevPage);
}

function NextPage() {
  let nextPage = data.current_page + 1;
  GetItems(data.folder, nextPage);
}

</script>

<template>
  <main class="parent">
    <sidenav-header />
    <top-header @previous-page="PreviousPage" @next-page="NextPage" @refresh-email="refreshItems"
      @compose-email="ComposeEmail" :title="data.folder" :current_page="data.current_page" />
    <item-list @item-selected="ItemSelected" :folder="data?.folders[data.folder]" :folderName="data.folder" />
    <sidenav @folder-selected="ChangeFolder" />
    <content :email="data.current_item" />
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
