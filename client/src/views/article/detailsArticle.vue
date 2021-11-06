<template>
    <div>
        <keep-alive>
            <router-view v-if="$route.meta.keepAlive" />
        </keep-alive>
            <router-view v-if="!$route.meta.keepAlive"></router-view>
            <WebBottom></WebBottom>
    </div>
</template>

<script>
import WebBottom from "@/components/common/WebBottom.vue"
    export default {
        name:"detailsArticle",
        components: {
            WebBottom
        }
    }
</script>

<style scoped>

</style>