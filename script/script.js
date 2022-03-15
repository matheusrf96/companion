// c = window
// m = document
// p = 'HASH'

(function(c,m,p){
c.cmp=c.cmp||function(){(c.cmp.q=c.cmp.q||[]).push(arguments)};
var i=m.createElement('input');i.type='hidden';i.id='eh';i.value=p;
var o=m.createElement('script');o.async=1;o.src='http://localhost:8080/cmp';
var n=m.getElementsByTagName('script')[0];n.parentNode.insertBefore(o, n);})
(window,document,'HASH');