<template>
<div>
   <v-row justify="space-around" align="center" style="margin: 5px">
        <v-col cols="6" align="center"  align-self="center">
          <h2>Light On:</h2>
          <h3>Currently: {{current_time.time_on}}</h3>
          <v-time-picker width="230px" v-model="start" format="24hr"></v-time-picker>
        </v-col>
        <v-col cols="6" align="center"  align-self="center">
          <h2>Light Off: </h2>
          <h3>Currently: {{current_time.time_off}}</h3>
          <v-time-picker width="230px" v-model="end" format="24hr"></v-time-picker>
        </v-col>
      </v-row>
<v-row justify="space-around" align="center" style="margin: 5px 0px 5px 0px">
     <v-col cols="6" align="center"  align-self="center">
     <v-btn id="button"   :disabled="TimeDisable" @click="sendLightTimes()" rounded color="accent" height="50" width="360">
           Save
            <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
          </v-btn>
     </v-col>
    <v-col cols="6" align="center"  align-self="center">
     <v-switch large v-model="current_time.state" @change="changeLight()">
        <template v-slot:label>
        <h3>Light Switch</h3>
        </template>
      </v-switch>
    </v-col>
</v-row>
</div>
</template>

<script>
//style="width: 290px; flex: 0 1 auto;"
    import axios from "axios"
    export default {
       
        name: "TimeSettings",
        components:{
            
        },
          data(){
             return{
            lightStatus: false,
            start: null,
            end: null,
            current_time: {
                    time_on: null,
                    time_off: null,
                    state: null
                },  
             };
         },
         methods: {
              getLightTimes: function () {
                axios.get("http://127.0.0.1:3000/adminSettings/get-light")
                    .then(result => {
                        this.current_time.time_on = result.data.time_on
                        this.current_time.time_off = result.data.time_off
                        this.current_time.state = !!result.data.state
                        /*if (result.data.state == 1){
                            this.current_time.state =  true
                        }else{
                            this.current_time.state  = false
                        }
                        */
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },

            sendLightTimes:function () {
                axios.post("http://127.0.0.1:3000/adminSettings/insert-light",
                    {time_on: this.start, time_off: this.end},
                    "content-type: application/json")
                    .then(this.current_time.time_on = this.start, this.current_time.time_off = this.end)
                    .catch(error => {
                        console.log(error);
                    });

            },
             changeLight:function () {
                console.log(this.time_on)
                axios.post("http://127.0.0.1:3000/adminSettings/changelight",
                    {state: this.current_time.state| 0},
                    "content-type: application/json")
                    .then()
                    .catch(error => {
                        console.log(error);
                    }); 

            },

         },
         created() {
             this.getLightTimes()
         },
         
         computed: {
         TimeDisable: function() {
            if((this.start == null) || (this.end == null)){
            return true
            }
            else{
            return false
            }
        }
        


         },
    };
</script>

<style scoped>
    #logo-text{
        font-family: Montserrat, sans-serif;
        font-style: normal;
        font-weight: 600;
        font-size: 36px;
        color: var(--v-primary-base)
    }
</style>