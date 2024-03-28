<template>
    <div
      class="ip-address"
      @focusin="showMessage = false"
      @focusout="isAddressValid"
    >
      <!-- Заголовок -->
      <FieldLabel :label="label" />
  
      <!-- Октет -->
      <div
        v-if="!readonly || (readonly && octets[0].length)"
        class="ip-address__wrapper"
        :class="{ 'ip-address__wrapper-readonly': readonly }"
      >
        <div
          v-for="(octet, idx) in octets"
          :key="idx"
          class="ip-address__control"
          :class="{
            'ip-address__control-dot': idx !== octets.length - 1 && !readonly,
            'ip-address__control-slash':
              idx === octets.length - 2 && showMask && !readonly,
          }"
        >
          <!-- Поле ввода -->
          <input
            v-if="!readonly"
            ref="refs"
            :value="octet"
            class="ip-address__control-input"
            :class="[{ 'ip-address-control_error': showMessage }, getRequired]"
            type="number"
            min="0"
            max="255"
            inputmode="numeric"
            @keydown="onSwitchArrow($event, idx)"
            @input="updateHandler($event, idx)"
            @paste="pasteHandler"
          />
  
          <span
            v-else
            class="ip-address__value"
            :class="{
              'ip-address__value-dot': idx !== octets.length - 1,
              'ip-address__value-slash': idx === octets.length - 2 && showMask,
            }"
          >
            {{ octet }}
          </span>
        </div>
  
        <SlideTransition orientation="x">
          <SvgCross v-if="isDataFilled && !readonly" @click="clearHandler" />
        </SlideTransition>
      </div>
  
      <span v-else class="ip-address__value">-</span>
  
      <!-- Сообщение об ошибке -->
      <ValidateTransition>
        <Message v-if="showMessage" :text="message" />
      </ValidateTransition>
    </div>
  </template>
  
  <script>
  import { computed, ref, watch } from 'vue'
  
  /** @module Validators - Проверка на соответствие заполнения IP адреса */
  import { regexpIPAddress } from '@/components/ui/validations/ip/ip.js'
  
  /** @module Components - Компоненты */
  import FieldLabel from '@/components/ui/components/FieldLabel.vue'
  import Message from '@/components/ui/components/Message.vue'
  
  /** @module Icons - Иконки */
  import SvgCross from '@/components/ui/icons/SvgCross.vue'
  
  /** @module Transitions - Переходы */
  import SlideTransition from '@/components/ui/transitions/SlideTransition.vue'
  import ValidateTransition from '@/components/ui/transitions/ValidateTransition.vue'
  
  export default {
    name: 'IPAddress',
  
    components: {
      SlideTransition,
      SvgCross,
      ValidateTransition,
      Message,
      FieldLabel,
    },
  
    props: {
      /** @param {string} label - Заголовок */
      label: {
        type: String,
        default: '',
      },
  
      /** @param {string} message - Текст ошибки
       * Example: text: {message: "hello", status: 0}
       * @param {number} status - Статус сообщения
       * 0 - Ошибка
       * 1 - Предупреждение
       * 2 - Успех */
      message: {
        type: Object,
        default: () => {},
      },
  
      /** @param {string} modelValue - Значение */
      modelValue: {
        type: String,
        default: '',
      },
  
      /** @param {boolean} readonly - Только ли для чтения */
      readonly: {
        type: Boolean,
        default: false,
      },
  
      /** @param {boolean} showMask - Отображать ли октет для маски подсети */
      showMask: {
        type: Boolean,
        default: false,
      },
  
      /** @param {required} - Обязательно ли поле для заполнения */
      required: {
        type: Boolean,
        default: true,
      },
    },
  
    emits: ['update:modelValue'],
  
    setup(props, { emit }) {
      const refs = ref([])
  
      const showMessage = ref(false)
  
      const octets = ref([])
  
      /** @computed
       * @name isDataFilled - Валидация заполненности данных */
      const isDataFilled = computed(() => {
        return octets.value.every((octet) => octet.length > 0)
      })
  
      /** @computed
       * @name convertToString - Конвертация массива объектов в строку */
      const convertToString = computed(() => {
        const array = octets.value.map((octet) => octet)
  
        return isDataFilled.value ? array.join('.') : array.join('')
      })
  
      /** @computed
       * @name getRequired - Является ли поле обязательным для заполнения */
      const getRequired = computed(() => {
        return props.required &&
          !props.readonly &&
          octets.value.every((octet) => !octet)
          ? 'app-field__required'
          : ''
      })
  
      /** @function
       * @name isAddressValid - Валидация корректности ввода данных */
      const isAddressValid = () => {
        showMessage.value = !isDataFilled.value
  
        if (props.modelValue === convertToString.value) return
  
        if (!showMessage.value) emit('update:modelValue', convertToString.value)
      }
  
      /** @function
       * @name onSwitchArrow - Переключение фокуса между полями ввода клавишами с клавиатуры
       * @param {string} key - Ключ нажатой клавиши
       * @param {number} idx - Порядковый номер элемента */
      const onSwitchArrow = ({ key }, idx) => {
        const rules = { ArrowLeft: idx - 1, ArrowRight: idx + 1 }
  
        if (Object.hasOwn(rules, key) && !!refs.value[rules[key]]) {
          refs.value[rules[key]].focus()
        }
      }
  
      /** @function
       * @name clearHandler - Очистка значения */
      const clearHandler = () => {
        octets.value = octets.value.map(() => '')
        emit('update:modelValue', convertToString.value)
      }
  
      /** @function
       * @name pasteHandler - Вставка значения из буфера обмена
       * @param {ClipboardEvent} event - Объект события вставки */
      const pasteHandler = (event) => {
        const text = event.clipboardData.getData('text')
  
        if (regexpIPAddress.test(text)) {
          text.split('.').forEach((el, idx) => (octets.value[idx] = el ?? ''))
        } else {
          showMessage.value = true
        }
  
        event.preventDefault()
      }
  
      /** @function
       * @name updateHandler - Обновление значения поля
       * @param {HTMLElement} target - Активный октет
       * @param {number} idx - Порядковый номер октета */
      const updateHandler = ({ target }, idx) => {
        const maxOctetLength = 3
  
        target.value = target.value.replace(/\D/gi, '')
  
        if (target.value[0] === '0') {
          if (!idx) {
            target.value = target.value.replace(/^0/gi, '')
          } else {
            if (/^00/gi.test(target.value) && target.value.length > 1) {
              target.value = 0
            } else if (/^0/gi.test(target.value) && target.value.length === 1) {
              target.value = 0
            } else {
              target.value = target.value.replace(/^0/gi, '')
            }
  
            if (idx < octets.value.length - 1) {
              refs.value[idx + 1].focus()
            }
          }
        }
  
        if (target.value > 255) target.value = 255
  
        octets.value[idx] = target.value
  
        if (!target.value && !!idx) return refs.value[idx - 1].focus()
  
        if (
          target.value.length === maxOctetLength &&
          idx < octets.value.length - 1
        ) {
          refs.value[idx + 1].focus()
        }
      }
  
      if (props.showMask) {
        octets.value = ['', '', '', '', '']
      } else {
        octets.value = ['', '', '', '']
      }
  
      /** @watch
       * @name isModelValueUpdated - Значение было изменено */
      watch(
        () => props.modelValue,
        (val) => {
          if (val) {
            octets.value = props.modelValue.split('.')
          }
        },
        { immediate: true }
      )
  
      return {
        refs,
        octets,
        showMessage,
        isDataFilled,
        getRequired,
        updateHandler,
        onSwitchArrow,
        pasteHandler,
        clearHandler,
        isAddressValid,
      }
    },
  }
  </script>