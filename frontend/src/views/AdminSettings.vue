<template>
  <div id="Admin">
    <AdminTopRow v-bind:headtext="'Settings'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
  <div style="margin: 40px 50px 0 50px;">
    <v-expansion-panels class="mb-6">
      <v-expansion-panel>
        <v-expansion-panel-header>
          E-Mail Settings
          <template v-slot:actions>
            <v-icon color="primary" large>mdi-email-edit</v-icon>
          </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
 <h1>E-Mail</h1>
        </v-expansion-panel-content>
      </v-expansion-panel>
      <v-expansion-panel>
        <v-expansion-panel-header disable-icon-rotate>

          Pump Settings
         
          <template v-slot:actions>
            <v-icon color="blue" large>mdi-water-pump</v-icon>
          </template>

        </v-expansion-panel-header>
        <v-expansion-panel-content>
<PumpSettings></PumpSettings>
        </v-expansion-panel-content>
      </v-expansion-panel>

      <v-expansion-panel>
        <v-expansion-panel-header disable-icon-rotate>
          Light Settings
          <template v-slot:actions>
            <v-icon color="yellow" large>mdi-lightbulb-on</v-icon>
          </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
<TimeSettings></TimeSettings>
        </v-expansion-panel-content>
      </v-expansion-panel>


      <v-expansion-panel>
        <v-expansion-panel-header disable-icon-rotate>
          Pin Settings
          <template v-slot:actions>
            <v-icon color="grey" large>mdi-lock-question</v-icon>
          </template>
        </v-expansion-panel-header>
        <v-expansion-panel-content>
            Pin Settings
        </v-expansion-panel-content>
      </v-expansion-panel>
    </v-expansion-panels>
  </div>
  </div>
</template>


<script>
import AdminTopRow from "@/components/admin/AdminTopRow.vue";
import TimeSettings from "@/components/admin/TimeSettings.vue";
import PumpSettings from "@/components/admin/PumpSettings.vue";

export default {
  name: "AdminSettings",
  components: {
    AdminTopRow,
    TimeSettings,
    PumpSettings
  },
  data() {
    return {};
  },
  methods: {}
};
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

span {
  color: var(--v-primary-base);
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
}

.no-hover:hover {
  background-color: transparent;
  text-decoration: none;
}

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}
</style>