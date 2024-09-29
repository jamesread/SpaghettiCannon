<template>
  <input @keypress="onKeyDown" placeholder = "Enter weight update here" />
</template>

<script language = "ts">
  export default { 
    setup() {
      const onKeyDown = (e) => {
        console.log(e.target.value)

        if (e.keyCode === 13) {
          sendUpdate()

        }
      }

      return {
        onKeyDown
      }
    }
  }
</script>
