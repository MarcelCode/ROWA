import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import vuetify from "vuetify"
import Vuex, { mapState } from 'vuex'
import InfoBoxPlants from '../src/components/home/InfoBoxPlants.vue'
import Farming from '../src/views/Farming.vue'
import axios from 'axios';


const plant = ({t,n})=>({
    plant_type: t,
    available_plants:n
})
jest.mock('axios');
describe('Farming',  () => {
  let wrapper;
  let localVue;
  beforeEach(() => {
    localVue =  createLocalVue()
    localVue.use(vuetify)
    localVue.use(Vuex)
    axios.get.mockResolvedValue({})
    wrapper = shallowMount(Farming, {localVue})
  })

  it('has a created hook', async()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })

  it('contains the right components', () => {
    expect(wrapper.contains(InfoBoxPlants)).toBe(true)

  })

  it('displays the correct amount of harvestable plants', async() => {    
    let data = {
      data:  [plant({t:'Test',n:5}), plant({t:'Flower', n:1})]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(Farming)
    await wrapper.vm.$nextTick()
    expect(wrapper.vm.harvestable_plants[1].available_plants).toBe(1)
    expect(wrapper.vm.harvestable_plants[0].plant_type).toBe('Test')
    expect(wrapper.vm.harvestable_plants.length).toBe(2)
  })

  it('displays the correct amount of plantable plants', async() => {
    let data = {
      data:  [plant({t:'Test',n:5})]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(Farming)
    await wrapper.vm.$nextTick()
    expect(wrapper.vm.plantable_plants[0].available_plants).toBe(5)
    expect(wrapper.vm.plantable_plants[0].plant_type).toBe('Test')
    expect(wrapper.vm.plantable_plants.length).toBe(1)
  })

  
})