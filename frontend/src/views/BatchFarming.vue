<template>
  <div id="BatchFarming">
    <AdminTopRow v-bind:headtext="'Batch Farming'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
  <v-row justify="center" align="center">    <BatchButtonPill></BatchButtonPill>  </v-row>
 <v-row justify="center" align="center"> <p>Plant or harvest multiple plants at once from this screen. Simply select which plants you planted or harvested</p>
 </v-row>
     <div v-if="batch_active">
       
                 <div  v-for="item in harvestableModules" :key="item">
       
   <v-row class="info-box" justify="center" align-content="center" align="center"> 
     
       <v-col align-self="center" align="center"> <h4>Plant-Type: </h4><br> <h5>{{item.plant_type}}</h5>  </v-col>
       <v-col align-self="center" align="center"> <h4> Module:</h4> <h5>{{item.module_position}} </h5><h4>Position: </h4><h5>{{item.plant_position}}</h5>   </v-col>
        <v-col > <img :src="getImgUrl(item.plant_type)" alt="" width="120px" height="auto"> </v-col>
        <v-col > <v-checkbox v-model="selectedHarvestable" label="Harvested?" :value="item"></v-checkbox> </v-col>
     
   </v-row>
   
    </div>
    <v-row justify="center" style="margin-top: 40px"> 
                <v-btn id="button" rounded color="accent" height="50" width="360" @click="sendHarvesting()">Save 
                  <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
                </v-btn>
    </v-row>
            </div>
            <div v-else>

              <div  v-for="item in plantableModules" :key="item">
       
   <v-row class="info-box" justify="center"> 
     
       <v-col align-self="center" align="center"> <h4>Plant-Type: </h4><h5>{{item.plant_type}} </h5>  </v-col>
       <v-col align-self="center" align="center"> <h4> Module: </h4> <h5> {{item.planted_module}} </h5>  </v-col>
        <v-col> <img :src="getImgUrl(item.plant_type)" alt="" width="120px" height="auto"> </v-col>
         <v-col> <v-checkbox v-model="selectedPlantable" label="Planted?" :value="item.planted_module"></v-checkbox> </v-col>
     
   </v-row>
   
    </div>
    <v-row justify="center" style="margin-top: 40px"> 
                <v-btn id="button" rounded color="accent" height="50" width="360" @click="sendPlanting()">Save 
                  <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
                </v-btn>
    </v-row>
            </div>
  </div>
</template>


<script>
import AdminTopRow from "@/components/admin/AdminTopRow.vue";
import BatchButtonPill from "@/components/admin/BatchButtonPill.vue";
import axios from "axios";
import {mapState} from "vuex"
export default {
  name: "BatchFarming",
  components: {
    AdminTopRow,
    BatchButtonPill
  },
  data() {
    return {
      plantableModules: null,
      harvestableModules: null,
      selectedPlantable: [],
      selectedHarvestable: []
    };
  },
  methods: {
    getImgUrl(pic) {
      return require("@/assets/harvesting/plants/" + pic + ".png");
    },
    sendPlanting: function() {
      axios
        .post(
          "http://127.0.0.1:3000/plant/plant-all",
          {planted_modules: this.makeString(this.selectedPlantable) },
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
         setTimeout(this.getPlantable, 500);
    },
    sendHarvesting: function() {
      axios
        .post(
          "http://127.0.0.1:3000/harvest/harvest-all",
          this.selectedHarvestable,
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
        setTimeout(this.getHarvestable, 500);
        
    },
    getPlantable: function() {
      axios
        .get("http://127.0.0.1:3000/plant/get-all")
        .then(result => {
          this.plantableModules = result.data;
        })
        .catch(error => {
          console.log(error);
        });
    },
    getHarvestable: function() {
      axios
        .get("http://127.0.0.1:3000/harvest/get-all")
        .then(result => {
          this.harvestableModules = result.data;
         
        })
        .catch(error => {
          console.log(error);
        });
    },
    makeString: function(NumberArray){
      var NumberString = ","
      for (var number of NumberArray) {
          NumberString = NumberString + number.toString() + ","
    }
      return NumberString
    }
  },
  computed: {
          ...mapState(["batch_active"])
  },
  created() {
    this.getPlantable();
    this.getHarvestable();
  }
};
</script>

<style scoped>
p {
  text-align: center;
  width: 410px;
  color: #828282;
  padding-top: 10px;
  margin: 0 !important;
}
h1,
h2 {
  font-weight: normal;
}

h4{
  font-family: Raleway;
font-style: normal;
font-weight: 500;
font-size: 12px;

/* #828282 */

color: #828282;
}
h5{

font-family: Montserrat;
font-style: normal;
font-weight: 600;
font-size: 18px;

/* #789659 */

color: #789659;
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