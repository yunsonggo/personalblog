<template>
    <div>
        <router-view></router-view>
        <WebBottom></WebBottom>
    </div>
</template>

<script>
import WebBottom from "@/components/common/WebBottom.vue"
    export default {
        name:"home",
        components: {
            WebBottom
        }
    }
</script>

<style scoped>

</style>