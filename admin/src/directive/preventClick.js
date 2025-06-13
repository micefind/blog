export default {
  mounted(el, binding) {
    el.addEventListener('click', () => {
      if (!el.disabled) {
        el.disabled = true

        setTimeout(() => {
          el.disabled = false
        }, binding.value || 1500) // 默认 1500ms 后可再次点击
      }
    })
  },
}
