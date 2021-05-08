var config = require('../../config/config.js');

Page({
    data: {
        id           : '',
        product      : null,
        swiperHeight : '',
    },
    onLoad: function(options) {
        var self = this;
        this.setData({
            id: options.id
        });

        wx.getSystemInfo({
            success: function(res) {
                self.setData({
                    swiperHeight : (res.screenWidth) + 'px'
                });
            }
        });

        wx.request({
            url: config.api.reqProductDetail.replace(':id', options.id),
            success: function(res) {
                var product = res.data.data.product || null;
                product.image.url = config.static.imageDomain + product.image.url;
                product.price     = product.price.toFixed(2);
                for (var i = 0; i < product.images.length; i++) {
                    var url = product.images[i].url;
                    product.images[i].url = config.static.imageDomain + url;
                }
                product.contents = JSON.parse(product.detail);
                for (var i = 0; i < product.contents.length; i++) {
                    if (product.contents[i].type == 'image') {
                        var url = product.contents[i].value;
                        product.contents[i].value = config.static.imageDomain + url;
                    }
                }
                self.setData({
                    product: product 
                });
            }
        });
    }
})