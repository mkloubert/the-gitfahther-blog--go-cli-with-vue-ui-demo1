<template>
  <v-container class="fill-height">
    <v-responsive
      class="align-centerfill-height mx-auto"
      max-width="640"
    >
      <v-row>
        <v-col cols="12">
          {{ timeFromApi }}
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <v-btn variant="outlined" block @click="reloadTimeFromApi">Get time</v-btn>
        </v-col>
      </v-row>
    </v-responsive>
  </v-container>
</template>

<script lang="ts">
  // MIT License
  //
  // Copyright (c) 2024 Marcel Joachim Kloubert (https://marcel.coffee)
  //
  // Permission is hereby granted, free of charge, to any person obtaining a copy
  // of this software and associated documentation files (the "Software"), to deal
  // in the Software without restriction, including without limitation the rights
  // to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
  // copies of the Software, and to permit persons to whom the Software is
  // furnished to do so, subject to the following conditions:
  //
  // The above copyright notice and this permission notice shall be included in all
  // copies or substantial portions of the Software.
  //
  // THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
  // IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
  // FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
  // AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
  // LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
  // OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
  // SOFTWARE.

  import axios from 'axios';
  import dayjs from 'dayjs';

  export default {
    data: () => {
      return {
        isBusy: false,
        timeFromApi: null as (string | null)
      };
    },
    methods: {
      reloadTimeFromApi: async function() {
        this.isBusy = true;
        
        try {
          // access the GET endpoint in
          // Go backend
          const {
            data,
            status
          } = await axios.get('/api/times/now', {
            responseType: 'text'
          });

          if (status !== 200) {
            throw new Error(`Unexpected response: ${status}`);
          }

          // write string to `timeFromApi`
          // so it can be print out in the GUI
          this.timeFromApi = dayjs(data).toISOString();
        } catch (error) {
          console.error('[ERROR]', 'HelloWorld.reloadTimeFromApi()', error);
        } finally {
          this.isBusy = false;
        }
      }
    },
    mounted() {
        this.reloadTimeFromApi();
    },
  }
</script>
