"use strict";

import { createApp, defineAsyncComponent } from "vue";
window.createApp = createApp;
window.defineAsyncComponent = defineAsyncComponent;

import axios from "axios";
window.axios = axios;