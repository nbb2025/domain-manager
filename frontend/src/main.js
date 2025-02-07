import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { Button, Card, Form, Input, Alert, Menu, Layout, Table, Modal, Select, Space, Tag } from '@arco-design/web-vue'
import App from './App.vue'
import router from './router'
import './assets/main.css'

const app = createApp(App)

app.use(createPinia())
app.use(router)

// 按需注册组件
app.use(Button)
app.use(Card)
app.use(Form)
app.use(Input)
app.use(Alert)
app.use(Menu)
app.use(Layout)
app.use(Table)
app.use(Modal)
app.use(Select)
app.use(Space)
app.use(Tag)

app.mount('#app')