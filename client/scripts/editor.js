var editor = {
    canvas : undefined,
    context : undefined,
    buffer : [],
    point : { x : 2, y: 0 },
    lineHeight : 16,

    drawPoint : function () {
        this.context.fillRect (this.point.x * 8, this.point.y * (this.lineHeight), 1, this.lineHeight);
    },

    draw: function() {
        this.context.clearRect (0, 0, this.canvas.width, this.canvas.height);
        var i = 0;
        var y = 10;
        for (i = 0; i<this.buffer.length; ++i) {
            this.context.fillText (this.buffer[i], 0, y + i * this.lineHeight);
        }

        this.drawPoint();
    },

    keypressHandler: function(e) {
        switch (String.fromCharCode(e.charCode))
        {
            case 'h': this.point.x = Math.max (0, this.point.x - 1); break;
            case 'l': this.point.x = Math.min (this.buffer[this.point.y].length - 1, this.point.x + 1); break;
            case 'k': this.point.y = Math.max (0, this.point.y - 1); break;
            case 'j': this.point.y = Math.min (this.buffer.length - 1, this.point.y + 1); break;
        }
        this.draw();
    },

    init: function(canvs, contents) {
        this.canvas = canvs;
        this.context = this.canvas.getContext("2d");
        this.canvas.addEventListener ('keypress', function (e) { editor.keypressHandler(e); }, false);
        this.context.font = "16px monospace";
        this.buffer = contents.split ("\n");
        this.draw ();
    },
};

