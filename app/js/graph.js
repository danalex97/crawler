const jquery = require("jquery");
const d3 = require("d3");
const host = "http://127.0.0.1:30000/";

let graph = {};

function getTopologyInfo() {
  jquery.ajax({
    type: "GET",
    url: host,
    success: function (data) {
      data = JSON.parse(data.replace(/\'/g, "\""));
      let graph = {};

      graph.links = data;
      graph.nodes = [];
      let page_to_id = {}
      let idx = 0

      function addNode(node) {
        if (!page_to_id[node]) {
          graph.nodes.push({
              "page" : node,
              "group" : idx
          });
          page_to_id[node] = idx;
          idx++;
        }
      }

      graph.links.forEach(function logArrayElements(element, index, array) {
        source = element.source;
        target = element.target;

        addNode(source);
        addNode(target);
      });

      graph.links = [];
      data.forEach(function logArrayElements(element, index, array) {
        source = element.source;
        target = element.target;
        if (source != target && page_to_id[target]) {
          graph.links.push({
            "source" : page_to_id[source],
            "target" : page_to_id[target],
            "value"  : 5
          })
        }
      });


      let height = 800;
      let width = 1278;

      var color = d3.scale.category20();

      let force = d3.layout.force()
          .linkDistance(400)
          //.linkStrength(30)
          .size([width, height]);

      let svg = d3.select("body").append("svg")
          .attr("width", width)
          .attr("height", height);

      force
          .nodes(graph.nodes)
          .links(graph.links)
          .start();

      let link = svg.selectAll(".link")
        .data(graph.links)
        .enter().append("line")
        .attr("class", "link")
        .style("stroke-width", function(d) { return Math.sqrt(d.value); });

      let node = svg.selectAll(".node")
        .data(graph.nodes)
        .enter().append("circle")
        .attr("class", "node")
        .attr("r", 5)
        .style("fill", function(d) { return color(d.group); })
        .call(force.drag);

      node.append("title")
      .text(function(d) { return d.page; });

      var tooltip = svg.selectAll('.node')
        .append('div')
        .attr('class', 'tooltip');

      tooltip.append('div')
        .attr('class', 'label');

      node.on('mouseover', function(d) {
        console.log(d.page);
        tooltip.select('.label').html(d.page);
        tooltip.style('display', 'block');
      });

      node.on('mouseout', function() {
        tooltip.style('display', 'none');
      });

      force.on("tick", function() {
        link.attr("x1", function(d) { return d.source.x; })
            .attr("y1", function(d) { return d.source.y; })
            .attr("x2", function(d) { return d.target.x; })
            .attr("y2", function(d) { return d.target.y; });

        node.attr("cx", function(d) { return d.x; })
            .attr("cy", function(d) { return d.y; });
        });
  }})
};
