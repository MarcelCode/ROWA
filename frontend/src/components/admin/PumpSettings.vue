<template>
  <div>
    <v-row justify="space-around" align="center" style="margin: 5px">
      <v-col cols="12" align="center" align-self="center">
        <h2>Pump On:</h2>
        <h3>Currently: {{time_on}}</h3>
        <v-time-picker v-model="start"  format="24hr" :landscape="$vuetify.breakpoint.mdAndUp" full-width></v-time-picker>
      </v-col>
      
    </v-row>
    <v-row justify="space-around" align="center" style="margin: 5px 0px 5px 0px">
        <v-card
      class="mx-auto"
      max-width="400"

       width="400"
    >

         <v-card-text>
        <v-row
          class="mb-4"
          justify="space-between"
        >
          <v-col class="text-left">
            <span
              class="display-3 font-weight-light"
              v-text="duration"
            ></span>
            <span class="subheading font-weight-light mr-1"> Minutes</span>
        
          </v-col>
          <v-col class="text-right">
             <v-btn
          id="button"
          :disabled="TimeDisable"
          @click="sendPumpTimes()"
          rounded
          color="accent"
          
        >
          Save
          <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
        </v-btn>
          </v-col>
        </v-row>
  
        <v-slider
          v-model="duration"
         :rules="rules"
          track-color="grey"
          always-dirty
          :min="min"
          :max="max"
        >
          <template v-slot:prepend>
            <v-icon
              :color="color"
              @click="decrement"
            >
              mdi-minus
            </v-icon>
          </template>
  
          <template v-slot:append>
            <v-icon
              :color="color"
              @click="increment"
            >
              mdi-plus
            </v-icon>
          </template>
        </v-slider>
      </v-card-text>
    </v-card>
    </v-row>
  </div>
</template>

<script>
//style="width: 290px; flex: 0 1 auto;"
import axios from "axios";
export default {
  name: "PumpSettings",
  components: {},
  data() {
    return {
      min: 1,
      max: 25,
      start: null,
      time_on: null,
      duration: null,
      rules: [
        v => v >= 5 || 'We recommend at least 5 Minutes',
        v => v <= 15 || 'Theres no benefit from running over 15 Minutes',
      ],
      
    };
  },
  methods: {
    decrement () {
      this.duration--
    },
    increment () {
      this.duration++
    },
    getPumpTimes: function() {
      axios
        .get("http://127.0.0.1:3000/adminSettings/get-pump")
        .then(result => {
          this.time_on = result.data.time_on;
          this.duration = result.data.duration;
          /*if (result.data.state == 1){
                            this.current_time.state =  true
                        }else{
                            this.current_time.state  = false
                        }
                        */
        })
        .catch(error => {
          console.log(error);
        });
    },

    sendPumpTimes: function() {
      axios
        .post(
          "http://127.0.0.1:3000/adminSettings/insert-pump",
          { time_on: this.start, duration: this.duration },
          "content-type: application/json"
        )
        .then(
          (this.time_on = this.start)
        )
        .catch(error => {
          console.log(error);
        });
    },
  /*  changeLight: function() {
      console.log(this.time_on);
      axios
        .post(
          "http://127.0.0.1:3000/adminSettings/changelight",
          { state: this.current_time.state | 0 },
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
    }*/
  },
  created() {
    this.getPumpTimes();
  },

  computed: {
    TimeDisable: function() {
      if (this.start == null) {
        return true;
      } else {
        return false;
      }
    }
  }
};
</script>

<style scoped>
#logo-text {
  font-family: Montserrat, sans-serif;
  font-style: normal;
  font-weight: 600;
  font-size: 36px;
  color: var(--v-primary-base);
}
</style>