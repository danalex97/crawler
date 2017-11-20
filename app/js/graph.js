const jquery = require("jquery");
const host = "http://127.0.0.1:30000/";

let svg = d3.select("#graph-container")
      .append("svg")
      .attr("width", "1278px")
      .attr("height", "800px")
      .call(d3.zoom().on("zoom", function () {
              svg.attr("transform", d3.event.transform)
      }))
      .append("g"),
    width = +svg.attr("width"),
    height = +svg.attr("height");

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
      graph.links.forEach(function logArrayElements(element, index, array) {
        source = element.source;
        target = element.target;
        if (graph.nodes.length == 0
          || graph.nodes[graph.nodes.length - 1] != source) {
            graph.nodes.push({
              "id" : source
            });
        }
      });
      graph.links = [];
      data.forEach(function logArrayElements(element, index, array) {
        source = element.source;
        target = element.target;
        graph.links.push({
          "source" : {"id" : source},
          "target" : {"id" : target},
          "value"  : 1
        })
      });
      console.log(graph.links[0])
      console.log(graph.nodes[0])
      let total = 0;
      let defs = svg.append("svg:defs");

      defs.append("svg:pattern")
          .attr("width", 1)
          .attr("height", 1)
          .attr("id", "server")
          .append("svg:image")
          .attr("xlink:href", "img/page.png")
          .attr("width", 100)
          .attr("height", 100)
          .attr("x", 12)
          .attr("y", 0);

      let simulation = d3.forceSimulation()
                      .force("link", d3.forceLink().id(function(d) {
                          return d.id;
                      }).distance(150))
                      .force("charge", d3.forceManyBody().strength(-5))
                      .force("center", d3.forceCenter(650, 400))
                      .force("collide", d3.forceCollide(100));

      let link = svg.selectAll(".link")
                  .data(graph.links, function(d) {return d.id;})
                  .enter().append("line").attr("class", "link");

      let node = svg.selectAll(".node")
                  .data(graph.nodes, function(d) {return d.id;})
                  .enter().append("g").attr("class", "node")

                  .call(d3.drag()
                      .on("start", dragstarted)
                      .on("drag", dragged)
                      .on("end", dragended));

      let tip = d3.tip()
              .attr('class', 'd3-tip')
              .offset([-10, 0])
              .html(function(d) {
                console.log(d);
                return "<strong style='color:red'>Link: </strong><span> " + d.id + " </span>";
              });

      svg.call(tip);

      node.append("circle")
          .attr("r", 60)
          .style("fill", "url(#server)");

      node.on("mouseover", function(d) { tip.show(d, this); })
          .on("mouseout", function(d) { tip.hide(); });

      node.append("title")
          .text(function(d) { return d.id; });

      simulation.nodes(graph.nodes)
          .on("tick", ticked);

      simulation.force("link")
          .links(graph.links);

      function ticked() {
          link
              .attr("x1", function(d) { return d.source.x; })
              .attr("y1", function(d) { return d.source.y; })
              .attr("x2", function(d) { return d.target.x; })
              .attr("y2", function(d) { return d.target.y; });

          node
              .attr("transform", function(d) {
                  return "translate(" + d.x + ", " + d.y + ")";
              });
      }

      function dragstarted(d) {
          if (!d3.event.active)  {
              simulation.alphaTarget(0.3).restart();
          }
      }

      function dragged(d) {
          d.fx = d3.event.x;
          d.fy = d3.event.y;
      }

      function dragended(d) {
          if (!d3.event.active) {
           simulation.alphaTarget(0);
          }
          d.fx = null;
          d.fy = null;
      }

    }
  })
};
