<template>
  <div class="about main-container">
    <v-data-table
      :headers="headers"
      :items="disks"
      :items-per-page="15"
    ></v-data-table>
  </div>
</template>

<script>
import { mapState, mapMutations, mapGetters } from "vuex";
import { getDisks } from "@/api/api";

export default {
  data() {
    return {
      headers: [
        { text: "Host", value: "Host" },
        { text: "Device", value: "Device" },
        { text: "Avail", value: "Avail" },
        { text: "Used", value: "Used" },
        { text: "Size", value: "Size" },
        { text: "Usage", value: "Usage" },
      ],
      disks: [],
    };
  },
  mounted() {
    this.updateDisks(this.target);
  },
  computed: {
    ...mapState(["target"]),
    ...mapGetters(["getTarget"]),
  },
  watch: {
    getTarget(target) {
      this.updateDisks(target);
    },
  },
  methods: {
    ...mapMutations(["setTarget"]),
    updateDisks(target) {
      const that = this;
      getDisks(target).then(function (res) {
        that.disks = res.data;
      });
    },
  },
};
</script>
