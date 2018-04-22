import { ipcRenderer } from 'electron';

export default {
  install(Vue) {
    Vue.prototype.$ipc = ipcRenderer;
  },
};
