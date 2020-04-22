<template>
  <div id="Admin">
    <AdminTopRow v-bind:headtext="'Change Plant Types'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
    <v-carousel dark height="700" hide-delimiter-background font-color="primary">
      <v-carousel-item dark v-for="item in plantTypes" :key="item.plant_type">
        <v-sheet color="secondary" height="25%">
          <v-row justify="center" margin="auto" padding="auto" class="carussel">
            <tr>
              <td>Module {{item.typemodule}}</td>
              <td>: {{item.typename}}</td>
            </tr>
          </v-row>
        </v-sheet>
        <v-sheet color="secondary" height="50%">
          <v-row justify="center" margin="auto" padding="auto">
            <img :src="getImgUrl(item.typename)" alt width="300px" height="auto" />
          </v-row>
        </v-sheet>
        <v-sheet color="secondary" height="25%">
          <v-row justify="center" margin="auto" padding="auto" class="carussel">
            <v-col cols="5">
              <v-select
                rounded
                v-model="item.typename"
                :items="Types"
                :menu-props="{ maxHeight: '400' }"
                label="Select"
                dense
                solo
                dark
              ></v-select>
            </v-col>
            <v-col cols="5">
              <v-btn @click="sendModuleChanges(item.typemodule, item.typename)">
                Save
                <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
              </v-btn>
            </v-col>
          </v-row>
        </v-sheet>
      </v-carousel-item>
    </v-carousel>
  </div>
</template>


<script>
import AdminTopRow from "@/components/admin/AdminTopRow.vue";
import axios from "axios";
export default {
  name: "ModuleType",
  components: {
    AdminTopRow
  },
  data() {
    return {
      plantTypes: null,
      Types: []
    };
  },
  methods: {
    getImgUrl(pic) {
      return require("@/assets/harvesting/plants/" + pic + ".png");
    },
    sendModuleChanges: function(Numbe, Plant) {
      axios
        .post(
          "http://127.0.0.1:3000/adminSettings/insertmodule-change",
          { typename: Plant, typemodule: Numbe },
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
    },
    getPlantTypes: function() {
      axios
        .get("http://127.0.0.1:3000/adminSettings/get-types")
        .then(result => {
          this.plantTypes = result.data;
        })
        .catch(error => {
          console.log(error);
        });
    },
    getKnownTypes: function() {
      axios
        .get("http://127.0.0.1:3000/adminSettings/get-knowntypes")
        .then(result => {
          result.data.forEach(element => {
            this.Types.push(element.type);
          });
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    this.getPlantTypes();
    this.getKnownTypes();
  }
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

tr {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 500;
  font-size: 32px;
  line-height: 29px;
  color: var(--v-primary-base);
}
span {
  color: var(--v-primary-base);
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
}

.carussel {
  color: var(--v-primary-base);
}

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}
</style>