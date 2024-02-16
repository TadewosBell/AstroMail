// import { writable } from 'svelte/store';
// // import type { dante } from '../../wailsjs/go/models';
// import { GetDevice } from '../../wailsjs/go/main/App';
// import { EventsOn } from '../../wailsjs/runtime/runtime';

// function createDevices() {
//     const { subscribe, update } = writable(new Array<dante.DanteDevice>());

//     EventsOn('DeviceNew', (device) => addDevice(device));
//     EventsOn('DeviceLost', (device) => setConnected(device, false));
//     EventsOn('DeviceReconn', (device) => setConnected(device, true));

//     async function addDevice(name) {
//         const newDev = await GetDevice(name);
//         update((devs) => {
//             devs.push(newDev);
//             return devs;
//         });
//     }

//     function setConnected(name, conn) {
//         update((devs) => {
//             devs.find((dev) => dev.name === name).connected = conn;
//             return devs;
//         });
//     }

//     return {
//         subscribe,
//     };
// }

// export const devices = createDevices();
