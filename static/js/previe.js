var previe = new Vue({
  el: ".image_pages",
  data: {
    imageSelected: null,
    previeImage: null
  },
  methods: {
    imagesSelect(event){
      this.imageSelected = event.target.files[0];
      console.log(event)
      this.onUploadImage(event)
    },
    onUploadImage(event){
      var formData = new FormData();
      formData.append('image', this.imageSelected, this.imageSelected.name);
      axios.post('/create_previe', formData)
        .then( res => {
          this.imageSelected = JSON.parse(res.data)
          var path = String(this.imageSelected.Previe[0].Path)
          path = path.slice(1);
          path ='url('+path+')';
          var block = event.path[1];
          console.log($(block), path);
          block = $(block).children('button');
          block.css("background-image", String(path));
        })
    }
  }
})