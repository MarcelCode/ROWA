import { createLocalVue, mount,shallowMount } from '@vue/test-utils'
import FarmInfo from '../src/components/home/FarmInfo.vue'
import FarmTransition from '../src/components/main/FarmTransition.vue'
import InfoBoxPlants from '../src/components/home/InfoBoxPlants.vue'
import vuetify from "vuetify"
import axios from 'axios';


jest.mock('axios');

const plant = ({t,n})=>({
    plant_type: t,
    available_plants:n
})

describe('FarmInfo',  () => {
  let wrapper;
  beforeEach(() => {
    const localVue =  createLocalVue()
    localVue.use(vuetify)
    axios.get.mockResolvedValue({}) 
    wrapper = shallowMount(FarmInfo, {localVue})   
  })

  
  it('has a created hook',()  => {
    expect(wrapper.isVueInstance()).toBe(true)  
  })

  it('contains the right components', () => {
    expect(wrapper.contains(FarmTransition)).toBe(true)
    expect(wrapper.contains(InfoBoxPlants)).toBe(true)
  })

  it('saves the correct amount of harvestable plants', async() => {    
    let data = {
      data:  [plant({t:'Test',n:5}), plant({t:'Flower', n:1})]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(FarmInfo)
    await wrapper.vm.$nextTick()
    expect(wrapper.vm.harvestable[1].available_plants).toBe(1)
    expect(wrapper.vm.harvestable[0].plant_type).toBe('Test')
    expect(wrapper.vm.harvestable.length).toBe(2)
  })

  it('saves the correct amount of plantable plants', async() => {
    let data = {
      data:  [plant({t:'Test',n:5})]
    }
    axios.get.mockResolvedValue(data)
    wrapper = shallowMount(FarmInfo)
    await wrapper.vm.$nextTick()
    expect(wrapper.vm.plantable[0].available_plants).toBe(5)
    expect(wrapper.vm.plantable[0].plant_type).toBe('Test')
    expect(wrapper.vm.plantable.length).toBe(1)
  })
})