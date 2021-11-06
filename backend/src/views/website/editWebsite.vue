<template>
  <div>
    {{websiteInfo}}
  </div>
</template>

<script>
export default {
  name: "editWebsite",
  data() {
    return {
        
    };
  },
  props: {
    websiteInfo: {
      type: Object,
      default: () => ({}),
    },
  },
  mounted() {
    
  },
  methods: {

  },
};
</script>

<style scoped></style>
