import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    farm_active: true,
    batch_active: true,
    farm_info: {
      "1": {
        type: "lettuce",
        pos: {
          "1": {
            age: 3,
            harvestable: true
          },
          "2": {
            age: 10,
            harvestable: true
          },
          "3": {
            age: 17,
            harvestable: true
          },
          "4": {
            age: 24,
            harvestable: true
          },
          "5": {
            age: 31,
            harvestable: false
          },
          "6": {
            age: 38,
            harvestable: false
          }
        }
      },
      "2": {
        type: "lettuce",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "3": {
        type: "basil",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "4": {
        type: "basil",
        pos: {
          1: {
            age: null,
            harvestable: null
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: false
          }
        }
      },
      "5": {
        type: "lettuce",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "6": {
        type: "lettuce",
        pos: {
          1: {
            age: null,
            harvestable: null
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: false
          }
        }
      },
    }
  },
  mutations: {
    CHANGE_DASH_STATE(state){
      state.farm_active = !state.farm_active
    },
    CHANGE_BATCH_STATE(state){
      state.batch_active = !state.batch_active
    },
    FarmUpdate(state,  data){
      console.log(data)
      state.farm_info[data.moduleNumber].type = data[0].plant_type.toLowerCase()
      for(var i = 0;i<6;i++){
        if(data[i]!== undefined){
          state.farm_info[data.moduleNumber].pos[i+1].harvestable = data[i].harvestable
          state.farm_info[data.moduleNumber].pos[i+1].age = data[i].age
          console.log(state.farm_info[data.moduleNumber].pos[i+1].harvestable)
        }
        else{
          state.farm_info[data.moduleNumber].pos[i+1].harvestable = null
          state.farm_info[data.moduleNumber].pos[i+1].age = 0
        }
        
      }
      

    }
  },
  actions: {
    change_dash_state({commit}){
      commit("CHANGE_DASH_STATE")
    },
    change_batch_state({commit}){
      commit("CHANGE_BATCH_STATE")
    }
  }
})
